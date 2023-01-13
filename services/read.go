package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/abe/erp.api/configs"
	"github.com/abe/erp.api/models"
)

func ConvertInt(txt string) int64 {
	result, _ := strconv.ParseInt(txt, 0, 64)
	return result
}

func ReadEDI(obj *models.DownloadMailBox, userID *string) (err error) {
	// fmt.Printf("REC: %s\n", obj.ID)
	file, err := os.Open(fmt.Sprintf("./public/%s", obj.FilePath))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initailize the DB
	db := configs.Store

	// Get Unit
	txtUnit := "BOX"
	txtPartType := "PART"
	isSync := true
	if obj.MailType.Factory.Title == "INJ" {
		txtUnit = "COIL"
		txtPartType = "WIRE"
		isSync = false
	}

	var unitData models.Unit
	if err := db.First(&unitData, "title", txtUnit).Error; err != nil {
		log.Fatal(err)
	}

	var itemTypeData models.ItemType
	if err := db.First(&itemTypeData, "title", txtPartType).Error; err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// check File Type
	if obj.MailType.Title == "RECEIVE" {
		for scanner.Scan() {
			txt := scanner.Text()
			receiveKey := strings.ReplaceAll((txt[4:(4 + 12)]), " ", "")
			// fmt.Printf("REC: %s\n", receiveKey)
			// Initailize Part Data
			PartNo := strings.ReplaceAll(txt[76:(76+25)], " ", "")
			PartName := strings.ReplaceAll(txt[101:(101+25)], " ", "")
			SlugPartNo := strings.ReplaceAll(PartNo, "-", "")

			part := models.Part{
				Slug:        SlugPartNo,
				Title:       PartNo,
				Description: PartName,
				IsActive:    true,
			}

			if err := db.FirstOrCreate(&part, &models.Part{Slug: SlugPartNo}).Error; err != nil {
				logData := models.SystemLogger{
					UserID:      userID,
					Title:       fmt.Sprintf("บันทึกข้อมูล %s", SlugPartNo),
					Description: fmt.Sprintf("%s", err),
					IsSuccess:   false,
				}
				db.Create(&logData)
			}

			// Initailize Ledger
			ledger := models.Ledger{
				AreaID:     &obj.MailBox.Area.ID,
				FactoryID:  &obj.MailType.Factory.ID,
				PartID:     &part.ID,
				ItemTypeID: &itemTypeData.ID,
				UnitID:     &unitData.ID,
			}
			db.FirstOrCreate(&ledger, &models.Ledger{
				AreaID:    &obj.MailBox.Area.ID,
				FactoryID: &obj.MailType.Factory.ID,
				PartID:    &part.ID,
			})

			ediReceive := models.GEDIReceive{
				Factory:          obj.MailType.Factory.Title,
				FacZone:          txt[4:(4 + 12)],
				ReceivingKey:     txt[4:(4 + 12)],
				PartNo:           PartNo,
				PartName:         PartName,
				Vendor:           obj.MailType.Factory.Title,
				Cd:               "02",
				Unit:             unitData.Title,
				Whs:              obj.MailType.Factory.Title,
				Tagrp:            "C",
				RecType:          "",
				PlanType:         "",
				RecID:            txt[0:4],
				Aetono:           txt[4:(4 + 12)],
				Aetodt:           txt[16:(16 + 10)],
				Aetctn:           ConvertInt(txt[26:(26 + 9)]),
				Aetfob:           ConvertInt(txt[35:(35 + 9)]),
				Aenewt:           ConvertInt(txt[44:(44 + 11)]),
				Aentun:           txt[55:(55 + 5)],
				Aegrwt:           ConvertInt(txt[60:(60 + 11)]),
				Aegwun:           txt[71:(71 + 5)],
				Aeypat:           txt[76:(76 + 25)],
				Aeedes:           txt[101:(101 + 25)],
				Aetdes:           txt[101:(101 + 25)],
				Aetarf:           0, //ConvertInt(txt[151:(151 + 10)]),
				Aestat:           0, //ConvertInt(txt[161:(161 + 10)]),
				Aebrnd:           0, //ConvertInt(txt[171:(171 + 10)]),
				Aertnt:           0, //ConvertInt(txt[181:(181 + 5)]),
				Aetrty:           0, //ConvertInt(txt[186:(186 + 5)]),
				Aesppm:           0, //ConvertInt(txt[191:(191 + 5)]),
				AeQty1:           0, //ConvertInt(txt[196:(196 + 9)]),
				AeQty2:           0, //ConvertInt(txt[205:(205 + 9)]),
				Aeuntp:           0, //ConvertInt(txt[214:(214 + 9)]),
				Aeamot:           0, //ConvertInt(txt[223:(223 + 11)]),
				Plnctn:           ConvertInt(txt[26:(26 + 9)]),
				PlnQty:           ConvertInt(txt[196:(196 + 9)]),
				Minimum:          0,
				Maximum:          0,
				Picshelfbin:      "PNON",
				Stkshelfbin:      "SNON",
				Ovsshelfbin:      "ONON",
				PicshelfbasicQty: 0,
				Outerpcs:         0,
				AllocateQty:      0,
			}

			var receiveType models.ReceiveType
			if err := db.First(&receiveType, "prefix", receiveKey[:3]).Error; err != nil {
				log.Fatal(err)
			}
			// Initailize ReceiveEnt
			dte, _ := time.Parse("02/01/2006", ediReceive.Aetodt)
			receiveEnt := models.Receive{
				DownloadID:    &obj.ID,
				ReceiveTypeID: &receiveType.ID,
				ReceiveDate:   dte,
				TransferOutNo: receiveKey,
				TexNo:         "-",
				IsSync:        true,
				IsActive:      true,
			}
			if err := db.FirstOrCreate(&receiveEnt, models.Receive{TransferOutNo: receiveKey}).Error; err != nil {
				log.Fatal(err)
			}

			receiveDetail := models.ReceiveDetail{
				ReceiveID: &receiveEnt.ID,
				LedgerID:  &ledger.ID,
				PlanQty:   ediReceive.PlnQty,
				PlanCtn:   ediReceive.Plnctn,
				IsActive:  true,
			}

			db.FirstOrCreate(&receiveDetail, models.ReceiveDetail{
				ReceiveID: &receiveEnt.ID,
				LedgerID:  &ledger.ID,
			})

			receiveDetail.PlanQty = ediReceive.PlnQty
			receiveDetail.PlanCtn = ediReceive.Plnctn
			db.Save(&receiveDetail)

			var countReceive []models.ReceiveDetail
			db.Where("receive_id=?", receiveEnt.ID).Find(&countReceive)
			ctn := 0
			for _, x := range countReceive {
				ctn += int(x.PlanCtn)
			}

			receiveEnt.Item = int64(len(countReceive))
			receiveEnt.PlanCtn = int64(ctn)

			// Disable Sync AW Data
			receiveEnt.IsSync = isSync
			db.Save(&receiveEnt)
			// fmt.Printf("REC: %s ETD: %s\n", receiveKey[:3], dte)
		}
		return
	}

	rnd := 1
	for scanner.Scan() {
		// line := scanner.Text()
		// etd := strings.ReplaceAll(line[28:(28+8)], " ", "")
		// Upddte, _ := time.Parse("20060102150405", strings.ReplaceAll(line[141:(141+14)], " ", ""))  //, "%Y%m%d%H%M%S"),
		// Updtime, _ := time.Parse("20060102150405", strings.ReplaceAll(line[141:(141+14)], " ", "")) // "%Y%m%d%H%M%S"),
		// EtdDte, _ := time.Parse("20060102", etd)
		// OrderMonth, _ := time.Parse("20060102", strings.ReplaceAll(line[118:(118+8)], " ", ""))
		rnd++
	}
	return
}
