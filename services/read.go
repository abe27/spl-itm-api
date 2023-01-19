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
	if result, err := strconv.ParseInt(txt, 0, 64); err == nil {
		return result
	}
	return 0
}

func ConvertFloat(i string) float64 {
	if n, err := strconv.ParseFloat(i, 64); err == nil {
		return n
	}
	return 0
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
	if obj.MailType.Factory.Title != "INJ" {
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

			// Initailize Ledger
			ledger := models.Ledger{
				WhsID:      &receiveType.WhsID,
				AreaID:     &obj.MailBox.Area.ID,
				FactoryID:  &obj.MailType.Factory.ID,
				PartID:     &part.ID,
				ItemTypeID: &itemTypeData.ID,
				UnitID:     &unitData.ID,
			}
			db.FirstOrCreate(&ledger, &models.Ledger{
				WhsID:     &receiveType.WhsID,
				AreaID:    &obj.MailBox.Area.ID,
				FactoryID: &obj.MailType.Factory.ID,
				PartID:    &part.ID,
			})

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
		line := scanner.Text()
		etd := strings.ReplaceAll(line[28:(28+8)], " ", "")
		Upddte, _ := time.Parse("20060102150405", strings.ReplaceAll(line[141:(141+14)], " ", ""))  //, "%Y%m%d%H%M%S"),
		Updtime, _ := time.Parse("20060102150405", strings.ReplaceAll(line[141:(141+14)], " ", "")) // "%Y%m%d%H%M%S"),
		EtdDte, _ := time.Parse("20060102", etd)
		OrderMonth, _ := time.Parse("20060102", strings.ReplaceAll(line[118:(118+8)], " ", ""))
		orderPlan := models.OrderPlan{
			Seq:              int64(rnd),
			Vendor:           obj.MailType.Factory.Title,
			Cd:               "",
			Sortg1:           "",
			Sortg2:           "",
			Sortg3:           "",
			Tagrp:            "C",
			PlanType:         "ORDERPLAN",
			Pono:             strings.ReplaceAll(line[13:(13+15)], " ", ""),
			RecId:            strings.ReplaceAll(line[0:4], " ", ""),
			Biac:             strings.ReplaceAll(line[5:(5+8)], " ", ""),
			EtdTap:           EtdDte, //), "%Y%m%d"),
			PartNo:           strings.ReplaceAll(line[36:(36+25)], " ", ""),
			PartName:         strings.TrimRight(line[61:(61+25)], " "),
			SampFlg:          strings.ReplaceAll(line[88:(88+1)], " ", ""),
			Orderorgi:        ConvertFloat(line[89:(89 + 9)]),
			Orderround:       ConvertFloat(line[98:(98 + 9)]),
			FirmFlg:          strings.ReplaceAll(line[107:(107+1)], " ", ""),
			ShippedFlg:       strings.ReplaceAll(line[108:(108+1)], " ", ""),
			ShippedQty:       ConvertFloat(line[109:(109 + 9)]),
			Ordermonth:       OrderMonth, //, "%Y%m%d" ),
			BalQty:           ConvertFloat(line[126:(126 + 9)]),
			Bidrfl:           strings.ReplaceAll(line[135:(135+1)], " ", ""),
			DeleteFlg:        strings.ReplaceAll(line[136:(136+1)], " ", ""),
			Reasoncd:         strings.ReplaceAll(line[138:(138+3)], " ", ""),
			Upddte:           Upddte,                                         //, "%Y%m%d%H%M%S"),
			Updtime:          Updtime,                                        // "%Y%m%d%H%M%S"),
			CarrierCode:      strings.ReplaceAll(line[155:(155+4)], " ", ""), //
			Bioabt:           ConvertInt(line[159:(159 + 1)]),
			Bicomd:           strings.ReplaceAll(line[160:(160+1)], " ", ""), //
			Bistdp:           ConvertFloat(line[165:(165 + 9)]),
			Binewt:           ConvertFloat(line[174:(174 + 9)]),
			Bigrwt:           ConvertFloat(line[183:(183 + 9)]),
			Bishpc:           strings.ReplaceAll(line[192:(192+8)], " ", ""), //
			Biivpx:           strings.ReplaceAll(line[200:(200+2)], " ", ""), //
			Bisafn:           strings.ReplaceAll(line[202:(202+6)], " ", ""), //
			Biwidt:           ConvertFloat(line[212:(212 + 4)]),
			Bihigh:           ConvertFloat(line[216:(216 + 4)]),
			Bileng:           ConvertFloat(line[208:(208 + 4)]),
			LotNo:            strings.ReplaceAll(line[220:], " ", ""),
			Minimum:          0,
			Maximum:          0,
			Picshelfbin:      "PNON",
			Stkshelfbin:      "SNON",
			Ovsshelfbin:      "ONON",
			PicshelfbasicQty: 0,
			OuterPcs:         0,
			AllocateQty:      0,
			CreatedAt:        Updtime,
			IsReviseError:    true,
		}

		orderPlan.DownloadID = &obj.ID

		var orderZone models.OrderZone
		db.Select("id,whs_id").Where("value=?", orderPlan.Bioabt).Where("factory_id=?", obj.MailType.Factory.ID).First(&orderZone)
		orderPlan.OrderZoneID = &orderZone.ID

		affcode := models.Affcode{
			Title:       orderPlan.Biac,
			Description: "-",
			IsActive:    true,
		}
		db.FirstOrCreate(&affcode, &models.Affcode{Title: orderPlan.Biac})
		orderPlan.AffcodeID = &affcode.ID

		customer := models.Customer{
			Title:       orderPlan.Bishpc,
			Description: orderPlan.Bisafn,
			IsActive:    true,
		}
		db.FirstOrCreate(&customer, &models.Customer{Title: orderPlan.Bishpc})
		orderPlan.CustomerID = &customer.ID

		// Create LastInvoice
		lastInv := models.LastInvoice{
			FactoryID: &obj.MailType.Factory.ID,
			AffcodeID: &affcode.ID,
			OnYear:    ConvertInt(etd[:4]),
			IsActive:  true,
		}
		db.FirstOrCreate(&lastInv, &models.LastInvoice{
			FactoryID: &obj.MailType.Factory.ID,
			AffcodeID: &affcode.ID,
			OnYear:    ConvertInt(etd[:4]),
		})

		/// Revise Type
		reviseTitle := "-"
		var reviseData models.ReviseOrder
		if len(orderPlan.Reasoncd) > 0 {
			reviseTitle = orderPlan.Reasoncd[:1]
		}
		if err := db.First(&reviseData, "title=?", reviseTitle).Error; err != nil {
			panic(err)
		}
		if reviseData.ID != "" {
			orderPlan.IsReviseError = false
			orderPlan.ReviseOrderID = &reviseData.ID
		}

		// Part
		part := models.Part{
			Slug:        strings.ReplaceAll(orderPlan.PartNo, "-", ""),
			Title:       orderPlan.PartNo,
			Description: orderPlan.PartName,
			IsActive:    true,
		}
		db.FirstOrCreate(&part, &models.Part{Slug: strings.ReplaceAll(orderPlan.PartNo, "-", "")})
		part.Description = orderPlan.PartName
		db.Save(&part)

		// Ledger
		ledger := models.Ledger{
			WhsID:       orderZone.WhsID,
			AreaID:      &obj.MailBox.Area.ID,
			FactoryID:   &obj.MailType.Factory.ID,
			PartID:      &part.ID,
			ItemTypeID:  &itemTypeData.ID,
			UnitID:      &unitData.ID,
			DimWidth:    0,
			DimLength:   0,
			DimHeight:   0,
			GrossWeight: 0,
			NetWeight:   0,
			Qty:         0,
			Ctn:         0,
			IsActive:    true,
		}

		db.FirstOrCreate(&ledger, &models.Ledger{
			AreaID:     &obj.MailBox.Area.ID,
			FactoryID:  &obj.MailType.Factory.ID,
			PartID:     &part.ID,
			ItemTypeID: &itemTypeData.ID,
			UnitID:     &unitData.ID,
		})

		ledger.DimWidth = orderPlan.Biwidt
		ledger.DimLength = orderPlan.Bileng
		ledger.DimHeight = orderPlan.Bihigh
		ledger.GrossWeight = (orderPlan.Bigrwt / 1000)
		ledger.NetWeight = (orderPlan.Binewt / 1000)
		ledger.Qty = orderPlan.Bistdp
		db.Save(&ledger)

		orderPlan.LedgerID = &ledger.ID
		var pc models.Pc
		db.First(&pc, "title=?", strings.ReplaceAll(line[86:(86+1)], " ", ""))
		orderPlan.PcID = &pc.ID

		var comm models.Commercial
		db.First(&comm, "title=?", strings.ReplaceAll(line[87:(87+1)], " ", ""))
		orderPlan.CommercialID = &comm.ID

		var orderTypeData models.OrderType
		db.First(&orderTypeData, &models.OrderType{Title: strings.ReplaceAll(line[137:(137+1)], " ", "")})
		orderPlan.OrderTypeID = &orderTypeData.ID

		var shipment models.Shipment
		db.First(&shipment, "prefix=?", strings.ReplaceAll(line[4:(4+1)], " ", ""))
		orderPlan.ShipmentID = &shipment.ID

		var sampleFlg models.SampleFlg
		db.First(&sampleFlg, "title=?", orderPlan.SampFlg)
		orderPlan.SampleFlgID = &sampleFlg.ID

		var orderGroup models.OrderGroup
		db.Preload("OrderGroupType").Where("affcode_id=?", affcode.ID).Where("customer_id=?", customer.ID).First(&orderGroup)
		txtOrderGroup := "N" // For not group order
		if orderGroup.ID != "" {
			txtOrderGroup = orderGroup.OrderGroupType.Title
		}
		switch txtOrderGroup {
		case "N":
			txtOrderGroup = "ALL"

		case "S":
			txtOrderGroup = orderPlan.Pono

		case "E":
			txtOrderGroup = orderPlan.Pono[len(orderPlan.Pono)-3:]
			var chkGroup models.OrderGroup
			db.Preload("OrderGroupType").Where("sub_order like ?", "%"+strings.TrimSpace(txtOrderGroup)+"%").Where("affcode_id=?", affcode.ID).Where("customer_id=?", customer.ID).First(&chkGroup)
			if chkGroup.ID == "" {
				txtOrderGroup = "ALL"
			}

		case "F":
			txtOrderGroup = orderPlan.Pono[:3]
			switch orderPlan.Pono[:1] {
			case "#", "@":
				txtOrderGroup = orderPlan.Pono[:4]
			}
			// case "NESC", "ICAM":
			// 	txtOrderGroup = orderPlan.Pono[:4]
			// }

			var chkGroup models.OrderGroup
			db.Preload("OrderGroupType").Where("sub_order like ?", "%"+strings.TrimSpace(txtOrderGroup)+"%").Where("affcode_id=?", affcode.ID).Where("customer_id=?", customer.ID).First(&chkGroup)
			if chkGroup.ID == "" {
				txtOrderGroup = "ALL"
			}
			if chkGroup.ID == "" {
				txtOrderGroup = "ALL"
			}
		}
		orderPlan.OrderGroup = txtOrderGroup
		if err := db.Create(&orderPlan).Error; err != nil {
			panic(err)
		}
		rnd++
	}
	return
}

func SubStringWire(txt string) (string, string, string) {
	var kinds string = ""
	var size string = ""
	var color string = ""
	lK := strings.Index(txt, " ")
	if lK >= 0 {
		txtKind := txt[:lK]
		lSize := txt[lK+1:]
		lS := strings.Index(lSize, " ")
		txtColor := lSize[lS+1:]
		kinds = txtKind
		size = lSize[:lS]
		color = txtColor
	}
	return kinds, size, color
}
