package sho_test

import (
	"fmt"

	"github.com/go-dedup/simhash"
	"github.com/go-dedup/simhash/sho"
)

// to show the full code in GoDoc
type dummy struct {
}

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example_output() {
	var docs = [][]byte{
		[]byte(" MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic"),
		[]byte(" FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic"),
		[]byte(" ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic"),
		[]byte(" 2013 Ford Mustang Premium LX Convertible Excellent condition, all options except navigation. New Tires, Rims and Brakes. Leather interior, heated seats & mirrors, xenon lamps etc. car is very well taken care off. text me at if… 66,000km | Automatic"),
		[]byte(" 2011 Ford Ranger FX4 Pickup Truck 2011 Ford Ranger FX4 pickup for sale. Truck is in immaculate condition. No rust. No dents. No scratches. Low km's. Purchased used several years ago and driven very little. Boxliner. Hitch. All FX4… 33,032km | Automatic"),
		[]byte(" 2012 Ford Focus SE Sedan Extremely good condition and well maintained. Mainly highway miles and very low considering the year of vehicle. Comes with PS, PB, power windows, power locks, power/heated mirrors, heated seats,… 61,000km | Automatic"),
		[]byte(" 2014 Ford Focus SE Sedan It's a great car. There is no issue with it. I am personally driving for 7 months and I am really happy about the conditions of car. If you are interesting wiht car,please call me. 37,635km | Automatic"),
		[]byte(" Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic"),
		[]byte(" Silver 2016 Ford Edge SEL SUV, Crossover for Lease Takeover Hi there, I have a 2016 Ford Edge in perfect condition for lease takeover. Bi-weekly payment is $263.00 tax included. Exterior color is silver and interior is black leathter. It has lease protection… 16,000km | Automatic"),
		[]byte(" 2010 Ford Focus SES, Black, Fully Loaded, Low Mileage There's a lein on the car and when sold, I will use the buyers money to pay off the car. The car is fully loaded, summer and winter tires & car matts. Car has a few minor exterior scratches. Received… 85,000km | Automatic"),
		[]byte(" 2016 Ford F-150 Pickup Truck Jason canopy 2016 5.1/2 Comes with 2 keys Inside carpet and light No damage Excellent condition Call 10,000km | Automatic"),
		[]byte(" Like New 2015 Ford F-150 XTR Pickup Truck 2015 Ford F-150 XTR Supercab, white with tan interior. Truck was lady owned and driven and meticulously maintained with only 27000 kms. The truck is the closest you will find to a brand new vehicle.… 27,000km | Automatic"),
		[]byte(" 2017 Ford Escape FWD 4dr SE SUV, Crossover – Lease Take over 2017 Ford Escape FWD 4dr SE SUV, Crossover – Lease Take over Lease takeover starting first two weeks of October 2017 with 24mths left. I have put a down payment of $1700 on this lease initially… 10,000km | Automatic"),
		[]byte(" 2014 Ford Escape SE Selling 2014 Ford Escape. 4 cylinder Eco boost. Back up camera, heated front seats, keyless entry with 2 key fobs. Recent car proof included. No accidents. $15,500 certified. 56,500km | Automatic"),
		[]byte(" 2014 Ford Focus SE Excellent condition. Like new. Comes with snow tires on steel factory rims and a set of tires on alloy rims. Economical to operate. Heated seats, air tilt, cruise, power mirrors, power windows, am fm… 28,000km | Automatic"),
		[]byte(" 2010 Ford F-150 SuperCrew Pickup Truck, 5.4L 4x4 2010 F-150 4x4 Supercrew XLT, 5-1/2ft box with spray in bed liner and soft tonneau cover, 5.4L V8 engine, XTR pkg, tow pkg with built in brake controller, Very clean. located in Orangeville. Call… 81,295km | Automatic"),
		[]byte(" 2013 Ford Focus Titanium - Low KMs! For Private Sale by Original Owner - 2013 Black Ford Focus Hatchback. Fully loaded! Titanium package includes: leather/heated seats, sunroof, navigation, bluetooth, rear-view camera, push button… 50,700km | Automatic"),
		[]byte(" 2014 Ford Black 2014 Ford F150 XLT SuperCrew 4x4 Mint Condition 32,000 KM 5.L v8 Towing Package Chrome Running Boards Rear View Camera Metal Tonneau Cover by Peragon Sprayed in rubber bed liner and… 32,000km | Automatic | CarProof"),
		[]byte(" 2015 Ford F-350 XLT Pickup Truck Selling my 2015 Ford XLT F-350, Reg-Cab, 8' box, 45,000Km. I factory ordered this truck. The only add it doesn't have is the step tail gate, Ford warranty until 2020, also Ford maintenance package,… 45,435km | Automatic"),
		[]byte(" 2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic"),
		[]byte(" 2013 Ford Fusion Sedan 2013 Ford Fusion SE Ecoboost only 65KMs, comes equipped with black leather interior, rear view camera, navigation system, touch screen display, sunroof, heated front seats, dual climate control,… 65,665km | Automatic"),
		[]byte(" 2014 Ford Fiesta SE Hatchback 9500 OBO In need of money, selling my 2014 Ford Fiesta SE with 57000km on it, all electric, alloy wheels, very little gas consumption, very spacious, automatic transmission, front wheel drive FWD, bluetooth… 57,000km | Automatic"),
		[]byte(" Trade 2016 Ford F-150 chrome bumpers for your black ones I'm looking to trade my chrome bumpers and grill of your black bumpers and grill. Please call or text 1,234km | Automatic"),
		[]byte(" MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic"),
		[]byte(" ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic"),
		[]byte(" FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic"),
		[]byte(" 2016 Ford Focus SE Sedan This is the fabled old lady's car. Bought new Nov. 2016 and driven only 2,000 klms. before she stopped driving for good. This car is brand new and very clean. Still has that new car smell. Please… 2,035km | Automatic"),
		[]byte(" 2013 Ford F-150 XLT/XTR 4x4 3.5 Ecobust 2013 Ford F-150 XLT/XTR 4x4 3.5 Ecobust with 90000 km on it I'm a first owner. Always serviced at Dixie Ford Mississauga. Still on the warranty for 100k km or 5 years. Factory Max Tow package.… 90,000km | Automatic"),
		[]byte(" 2012 Ford Fiesta Hatchback Selling my Ford fiesta, car runs in great conditions brand new breaks and front routers, excellent on consuming gas 58000 KM Smoke Free Safety is available contact 647 909 7578 Thanks 58,000km | Automatic"),
		[]byte(" 2014 Ford Focus Automatic MINT CONDITION I am selling my 2014 focus. Excellent condition, very well maintained. Only 65000km. Economic both on Highway and city. Bluetooth, traction control, heated seats and much more. Call me or text me at… 65,000km | Automatic"),
		[]byte(" 2014 Ford Focus SE Sedan Great car! No accidents! Low mileage! Winter tires in price! It runs perfectly, everything in a perfect condition - You should come and see it! This car is driven by my wife and need a bigger one. 25,734km | Automatic"),
		[]byte(" ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic"),
		[]byte(" 2014 Ford Mustang V6 Premium 2014 Ford Mustang with Pony package. 304 horsepower. leather power seat. Bluetooth. three drive model:comfort normal sport. 18 tires. two rims for you. the original rims are brand new in the box.… 41,500km | Automatic"),
		[]byte(" 2014 ford Mustang v6 premium 2014 Ford Mustang with Pony package. 304 horsepower. leather power seat. Bluetooth. three drive model:comfort normal sport. 18 tires. two rims for you. the original rims are brand new in the box.… 41,500km | Automatic"),
		[]byte(" 2012 Ford Fiesta SES Ford Fiesta SES red leather interior fully loaded. Excellent condition with very low mileage. Never winter driven. Made for city driving, really good on gas. 9 months of financing at 0% interest… 35,100km | Automatic"),
		[]byte(" 2014 Ford Mustang Premium V6 2014 Ford Mustang with Pony package. 304 horsepower. leather power seat. Bluetooth. three drive model:comfort normal sport. 18 tires. two rims for you. the original rims are brand new in the box.… 41,500km | Automatic"),
		[]byte(" 2014 Ford Fusion Titanium Sedan Very clean former salesman's car, mostly highway-driven. Safety certificate, all-season tires, just detailed. Fully-loaded with all the options Ford offered in 2014. Priced to move at $19,699 + HST 77,000km | Automatic"),
		[]byte(" 12 Ford Focus SE, Low 58km Only! Clean history, no accident Winter tire & Summer tire Just got marry, need a bigger car for family!!! Text please, Jeremy Thanks! 58,000km | Automatic"),
		[]byte(" 2017 Ford Mustang GT CALIFORNIA SPECIAL Coupe (2 door) 2017 MUSTANG GT COUPE CALIFORNIA SPECIAL GRABBER BLUE INCLUDED ON THIS VEHICLE EQUIPMENT GROUP 401A SINGLE CD W/TUNER/SDARS/HD 12 SPEAKER SYSTEM OPTIONAL EQUIPMENT 5.0L 4V TI-VCT V8 ENGINE 6-SPEED… 5,540km | Automatic"),
		[]byte(" Low Mileage 2014 Ford Focus SE - Sport Package Hardly used with less than 15 thousand kilometers. Bought new in 2014 and in perfect condition. Heated seats and winter tires (on rims) included. Please call or text if interested. 14,500km | Automatic"),
		[]byte(" 2014 ford fusion Wrote off my car and this is the extra stuff in my garage that I need to geto rid of. Weather tech floor mats front+ back Factory rubber floor mats front+ back Factory carpet floor mats front+ back… 90,000km | Automatic"),
		[]byte(" 2013 Ford Escape SUV, Crossover VERY LOW KMs (37,500) FORD ESCAPE FOR SALE. This car has been very well maintained over its life and is 100% ACCIDENT FREE. Leather interior, heated seats = a comfortable ride. Only blemish on it is… 37,700km | Automatic | CarProof"),
		[]byte(" 2014 Ford Fusion SE Sedan-Amazing Car! Eco boost 2.0L I4 GTDI Engine Upgraded 18 Alloy Rims Set of Winter Tires Included (new 2015) Remote Keyless Entry New weather tech mats included Sync Voice Activated System Navigation System Rear… 51,500km | Automatic"),
		[]byte(" 2016 Ford F-250 Lariat Pickup Truck Diesel 2016 f250. Mint shape just no longer want payments. Looking for someone to take over. Fully loaded, has every option possible. 6.7 powerstroke, dpf and egr removed. Will come with dpf and egr. Has an… 40,000km | Automatic"),
		[]byte(" 2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic"),
		[]byte(" 2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic"),
		[]byte(" FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic"),
		[]byte(" MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic"),
		[]byte(" ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic"),
		[]byte(" 2014 Ford E-250 Minivan, Van 2014 E250 Super Duty Regular Box 4.6L with A/C, bucket reclining seats, full cargo divider, side barn door, plywood floor, 3M protection and serviced regularly. Looking for some to assume lease or… 87,750km | Automatic"),
		[]byte(" 2013 Ford Focus SE Sedan 2013 December Ford Focus SE sedan for sale . Low kilometres... only 9000 kilometers. Looks like brand new. Good condition. Please contact me on 9,000km | Automatic"),
		[]byte(" 2017 Ford Mustang EcoBoost Premium Coupe (2 door) I am looking to sell my 2017 Ford Mustang EcoBoost Premium. It is Grabber Blue in color. Driven by a professional accountant. Bought it in September 2016. Package details included: Premium package… 9,800km | Automatic"),
		[]byte(" 2013 Ford Fiesta SE Sedan Beautiful light blue Fiesta SE lady driven, non-smoking. Garage stored in the winter with no accidents. Great on gas and easy to park. Lots of options including LED lights for safe driving at night,… 87,000km | Automatic"),
		[]byte(" 2012 Ford Mustang Premium Sport Coupe Selling my rare 2012 Club of America Edition Mustang. The car is fully loaded with all options available. Car is very well maintained, always stored in garage year round and comes with a separate set… 75,000km | Automatic"),
		[]byte(" 2012 Ford Mustang Premium Sport Coupe (2 door) Selling my rare 2012 Club of America Edition Mustang. The car is fully loaded with all options available. Car is very well maintained, always stored in garage year round and comes with a separate set… 75,000km | Automatic"),
		[]byte(" 2010 Ford Mustang Hi I'm selling my 2010 ford mustang mint condition Power windows Power locks A/c 2 keys Back up sensors Set of tires on rims Safety and emission Car proof is provided 55,254km | Automatic"),
		[]byte(" 2014 Ford Focus Hatchback Moving out of the country. Have to sell. Mint condition with lots of add-ons + new winter tires. SE Hatch with Sports package. only 27 months on the road. Includes maintenance package until March… 25,000km | Automatic"),
		[]byte(" 2016 Ford Mustang Convertible 2016 Mustang in beautiful shape clean and like new. Never winter driven. One owner-bought brand new. No accidents and tires are like new with 5000 kilometers. electric seat, mirrors, keyless,… 5,000km | Automatic"),
		[]byte(" Ford Focus 2011 SE for sale Blue 4 Door Sedan, mint condition, low mileage (under 50,000km) 49,000km | Automatic"),
		[]byte(" Ford F-150 EcoBoost 2014 crew cab with box liner and cover. Low kilometers. Back up camera and sensors. Satellite radio and CD player cold AC and rubber floor mats. Cloth seats. Six seater, remaining bumper to bumper… 26,000km | Automatic"),
		[]byte(" Lease takeover!! 2016 Ford Edge- fwd black on black leather Automatic keyless entry, automatic trunk, panoramic sunroof, push start, automatic start with key. 24 months left on lease with 14,300 kms. I am currently… 14,300km | Automatic"),
		[]byte(" 2012 Ford Focus SEL Hatchback for Sale - 1 owner Ford Focus Hatchback for sale by owner. Gently used with low mileage. Maintenance has been done up to date. Recent oil change and disc brake replacement. The car has been detailed and cleaned for the… 76,261km | Automatic"),
		[]byte(" 2014 Ford F-150 XL Pickup Truck sold by owner 2014 Ford F-150 XL Pickup Truck sold by owner, low mileage (64000km). New tires, cap, hitch, no accidents, great conditions. First come first buy. Please, contact me 64,000km | Automatic | CarProof"),
		[]byte(" Ford Mustang Very clean vehicle comes with summer and winters on rims all dealer serviced paper work and all rust proofed every year breaks are brand new to start off the summer warranty still available I have a… 55,000km | Automatic | CarProof"),
		[]byte(" 2012 Ford F-250 Super Duty XL Crew Cab Pickup Truck Commercial Truck Like New 2012 Ford F250 Super Duty XL Crew Cab 6.5' Box Grey 80K. Truck truly like Brand New & Very Clean and Powerful, it's barely used and had no accident whatsoever, also you'll… 80,000km | Automatic"),
		[]byte(" 2015 Ford F-150 Lariat SuperCrew Pickup Truck 2015 Ford F150 Lariat. Loaded. No sunroof. 500A Available August 1st, 2017 Presently in northern Ontario. Contact Martin 55,000km | Automatic"),
		[]byte(" 2015 Ford F-150 Lariat SuperCrew Pickup Truck 2015 Ford F150 Lariat. Loaded. No sunroof. 500A Available August 1st, 2017 Presently in northern Ontario. Contact Martin 55,000km | Automatic"),
		[]byte(" 2015 Ford F-150 SuperCrew FX4 Off Road sport Pickup Truck Mint condition 3.5 Eco Boost V6. Aluminum body. Still under warranty. Heated power Lombard seats. Brake pedal adjustment. Power windows, locks, mirrors, sunroof. Tilt and telescopic steering wheel.… 21,000km | Automatic"),
		[]byte(" Ford Fusion 2011 SE BLACK **CLEAN** 2011 Model Clean Ford Fusion SE Oil is newly changed. NO CRASH, ALL PARTS ORIGINAL. Only Left Mirror is Broken. (Still, works and can easily be replaced) AROUND 199.297 km If you want a test drive,… 199km | Automatic"),
		[]byte(" FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic"),
		[]byte(" MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic"),
		[]byte(" ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic"),
		[]byte(" 2016 Ford F-150 Lariat Pickup Truck 2.7L 501A w/FX4 Offroad Pky Vehicle Features: 2016 Magnetic Grey Ford F-150 SuperCrew 145 Wheelbase FX4 Off Road Package w/Skid Plates, 3.73 Ratio E-lock 2.7L EcoBoost Twin-Turbo Engine 6-Speed SelectShift Transmission 6500#… 19,999km | Automatic"),
		[]byte(" 2015 Ford Escape Titanium SUV, Crossover Next to new Ford Escape Titanium model. Sun roof, 4WD, full set of winter tires & rims, power liftgate, Bluetooth, heated mirrors, steering wheel audio controls, rear window defogger, rearview camera… 41,193km | Automatic"),
		[]byte(" 2013 Ford F-150 SuperCrew Pickup Truck Great Clean truck, 301A package, Brakes just done, power mirrors, Pedals, Windows, Seats Bluetooth and satellite radio Trailer brake Controller and Class III hitch 80000 km 6.5 foot box 5.0L V8 motor… 80,000km | Automatic"),
		[]byte(" 2015 Ford Focus Sedan ***Reduced***2015 Ford Focus fully loaded, leather interior, power sun roof, full warranty, full maintenance package included, back up camera, list goes on. Fuel is amazing. Will sell or you may take… 55,000km | Automatic"),
		[]byte(" Ford Mustang 2015 Ecoboost - Excellent Condition Ford Mustang Ecoboost - Excellent Condition 2015 Ford Mustang EcoBoost 2.3 L Automatic, 305 Horse Power, Absolute Beauty Keyless Entry, Remote Starter, Bluetooth, Steering Tilt, Reverse Camera, Power… 51,000km | Automatic"),
		[]byte(" 2013 Ford Escape SE Price is firm. Immaculate - no pets, no kids, non-smoker, female driven, I do not eat or drink in the car. Just got oil changed. I absolutely love this vehicle, just need to downsize to a car. 1.6L… 68,000km | Automatic"),
		[]byte(" MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic"),
		[]byte(" 2013 Ford F-450 King Ranch Pickup Truck FORD EXTENDED WARRANTY PREMIUM CARE PLAN - AIRLIFT REAR AIR SUSPENSION WITH IN CAB AIR CONTROLLER. THIS IS NIT A CHEAP AFTERMARKET PRODUCT.… 30,500km | Automatic"),
		[]byte(" XLT Beautiful SUV. Has XLT package. Very low Kms's on it and hardly used. Message for details. 25,000km | Automatic"),
		[]byte(" 2013 Ford Escape SE SUV, Crossover Selling a 2013 Ford Escape SE fully loaded still under warranty from Ford Canada. Need to sell the car ASAP as I have no place to keep at home. Contact: pagal_boi@hotmail.com 80,000km | Automatic"),
		[]byte(" -- URGENT-- 2010 Ford F-150 XL -- MINT -- BULLET PROOF V8 — Hello, I have my beautiful Grey Ford F150 XL for sale. This truck has been driven by a senior to the grocery store and back. This truck has never been smoked in and is very clean base model. It has a… 41,500km | Automatic"),
		[]byte(" Ford Escape Lease take over, Ford Escape 2017,1.5 L /ecoboost 41 months left, payment is 420CAN per month. Very clean like a new car 8,000km | Automatic"),
		[]byte(" 2014 Ford Escape SE SUV, Crossover 4X4, NAVIGATION, BACK UP CAMERA, BLUETOOTH, ROOF RAIL, SUNROOF. Rebuilders title. CALL 40,200km | Automatic"),
		[]byte(" 2010 Ford Focus SE LOW KM 2010 Ford Focus SE with only 79,000 KM Great little car, no longer need car as not commuting to school, lady driven, Second Owner. Features: (listed as many as I could think of, email any questions… 79,000km | Automatic"),
		[]byte(" 2014 Ford Escape Titanium SUV, Crossover 4WD TITANIUM, POWER HEATED LEATHER SEATS, PANORAMIC SUNROOF, PREMIUM AUDIO WITH BLUETOOTH PHONE CONNECTIVITY, REMOTE STARTER, BACK-UP CAMERA AND MUCH MORE. THIS PREMIUM ECOBOOST IS IN EXCELLENT SHAPE… 54,400km | Automatic"),
		[]byte(" 2014 Ford Escape SE SUV, Lease takeover $344.00 Lease takeover. 11 months remaining on lease. Return date : May 28, 2018 80,000 km lease allowance. Current km reading 49,500 $344.00 per month - plus tax Pearl White, SE, 4wd, panoramic sunroof,… 49,500km | Automatic"),
		[]byte(" Mint 2014 Ford F-250 New tires Motor is good Coolant is leaking Transmission is good Selling as is 330km | Automatic"),
		[]byte(" 2015 Ford Escape 4WD 4dr SE SUV, Crossover Dual front climate control, Power windows, Power door locks, Power drivers seat, Power mirrors, Cloth seats, Heated seats, Bucket seats, Split bench seat, 17 inch alloy wheels, All season tires,… 23,000km | Automatic"),
		[]byte(" Lease takeover - 2017 Mustang GT Premium w/ California Special Looking for someone to take over my lease, I am moving to the States and I cannot take this vehicle with me. Car was leased Oct/16 with 25000km annual milage. Contact me for more information. 8,950km | Automatic"),
		[]byte(" 2011 Ford Escape Hatchback Clean and reliable vehicle, never had any problems with it. Comes with winter tires 85,000km | Automatic"),
	}

	// Code starts

	oracle := sho.NewOracle()
	sh := simhash.NewSimhash()
	r := uint8(3)
	for _, d := range docs {
		hash := sh.GetSimhash(sh.NewWordFeatureSet(d))
		if oracle.Seen(hash, r) {
			fmt.Printf("=: Simhash of %x for '%s' ignored.\n", hash, d)
		} else {
			oracle.See(hash)
			fmt.Printf("+: Simhash of %x for '%s' added.\n", hash, d)
		}
	}

	fmt.Println("================")
	oracle = sho.NewOracle()
	r = uint8(8)
	for _, d := range docs {
		hash := sh.GetSimhash(sh.NewWordFeatureSet(d))
		if h, nd, seen := oracle.Find(hash, r); seen == true {
			fmt.Printf("=: Simhash of %x ignored for %x (%d).\n", hash, h, nd)
		} else {
			oracle.See(hash)
			fmt.Printf("+: Simhash of %x added.\n", hash)
		}
	}

	fmt.Println("================")
	oracle = sho.NewOracle()
	r = uint8(8)
	for _, d := range docs {
		hash := sh.GetSimhash(sh.NewWordFeatureSet(d))
		if n := oracle.Search(hash, r); len(n) > 0 {
			fmt.Printf("!: Similiar found for %x (:%v).\n", hash, n)
		}
	}

	// Code ends

	// Output:
	// +: Simhash of c833835fb8ef4733 for ' MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic' added.
	// +: Simhash of c7ad53ee7a9354e for ' FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic' added.
	// +: Simhash of 58b2c31ce3ef3abd for ' ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic' added.
	// +: Simhash of 5876d51ea5eb2f7c for ' 2013 Ford Mustang Premium LX Convertible Excellent condition, all options except navigation. New Tires, Rims and Brakes. Leather interior, heated seats & mirrors, xenon lamps etc. car is very well taken care off. text me at if… 66,000km | Automatic' added.
	// +: Simhash of d8fafe1ee3eb3e0d for ' 2011 Ford Ranger FX4 Pickup Truck 2011 Ford Ranger FX4 pickup for sale. Truck is in immaculate condition. No rust. No dents. No scratches. Low km's. Purchased used several years ago and driven very little. Boxliner. Hitch. All FX4… 33,032km | Automatic' added.
	// +: Simhash of 5976ee1ea6ef375e for ' 2012 Ford Focus SE Sedan Extremely good condition and well maintained. Mainly highway miles and very low considering the year of vehicle. Comes with PS, PB, power windows, power locks, power/heated mirrors, heated seats,… 61,000km | Automatic' added.
	// +: Simhash of 8873df0fe6eb36bc for ' 2014 Ford Focus SE Sedan It's a great car. There is no issue with it. I am personally driving for 7 months and I am really happy about the conditions of car. If you are interesting wiht car,please call me. 37,635km | Automatic' added.
	// +: Simhash of 1832c51ee6eb2e3e for ' Ford F-150. Lariat DO NOT BUY. Truck has been in the shop 50 days so far. It has had a vibration since day one and Ford cannot get rid of it. The have done everything possible to the underside of this truck and it is… 11,000km | Automatic' added.
	// +: Simhash of 872f71fe0ebba3a for ' Silver 2016 Ford Edge SEL SUV, Crossover for Lease Takeover Hi there, I have a 2016 Ford Edge in perfect condition for lease takeover. Bi-weekly payment is $263.00 tax included. Exterior color is silver and interior is black leathter. It has lease protection… 16,000km | Automatic' added.
	// +: Simhash of d8dcc1186ba977be for ' 2010 Ford Focus SES, Black, Fully Loaded, Low Mileage There's a lein on the car and when sold, I will use the buyers money to pay off the car. The car is fully loaded, summer and winter tires & car matts. Car has a few minor exterior scratches. Received… 85,000km | Automatic' added.
	// +: Simhash of 8d63bf7ee661affc for ' 2016 Ford F-150 Pickup Truck Jason canopy 2016 5.1/2 Comes with 2 keys Inside carpet and light No damage Excellent condition Call 10,000km | Automatic' added.
	// +: Simhash of d89abd1ceee92e7c for ' Like New 2015 Ford F-150 XTR Pickup Truck 2015 Ford F-150 XTR Supercab, white with tan interior. Truck was lady owned and driven and meticulously maintained with only 27000 kms. The truck is the closest you will find to a brand new vehicle.… 27,000km | Automatic' added.
	// +: Simhash of 9872bf7ee2e13425 for ' 2017 Ford Escape FWD 4dr SE SUV, Crossover – Lease Take over 2017 Ford Escape FWD 4dr SE SUV, Crossover – Lease Take over Lease takeover starting first two weeks of October 2017 with 24mths left. I have put a down payment of $1700 on this lease initially… 10,000km | Automatic' added.
	// +: Simhash of 4472bd7fe4eb154c for ' 2014 Ford Escape SE Selling 2014 Ford Escape. 4 cylinder Eco boost. Back up camera, heated front seats, keyless entry with 2 key fobs. Recent car proof included. No accidents. $15,500 certified. 56,500km | Automatic' added.
	// +: Simhash of 4872e50fa6eb3576 for ' 2014 Ford Focus SE Excellent condition. Like new. Comes with snow tires on steel factory rims and a set of tires on alloy rims. Economical to operate. Heated seats, air tilt, cruise, power mirrors, power windows, am fm… 28,000km | Automatic' added.
	// +: Simhash of 9932b71ea4a934bd for ' 2010 Ford F-150 SuperCrew Pickup Truck, 5.4L 4x4 2010 F-150 4x4 Supercrew XLT, 5-1/2ft box with spray in bed liner and soft tonneau cover, 5.4L V8 engine, XTR pkg, tow pkg with built in brake controller, Very clean. located in Orangeville. Call… 81,295km | Automatic' added.
	// +: Simhash of c1b2ef3fe1b95f0e for ' 2013 Ford Focus Titanium - Low KMs! For Private Sale by Original Owner - 2013 Black Ford Focus Hatchback. Fully loaded! Titanium package includes: leather/heated seats, sunroof, navigation, bluetooth, rear-view camera, push button… 50,700km | Automatic' added.
	// +: Simhash of 8f2f34ea4eb364f for ' 2014 Ford Black 2014 Ford F150 XLT SuperCrew 4x4 Mint Condition 32,000 KM 5.L v8 Towing Package Chrome Running Boards Rear View Camera Metal Tonneau Cover by Peragon Sprayed in rubber bed liner and… 32,000km | Automatic | CarProof' added.
	// +: Simhash of 897af11ee6e9146f for ' 2015 Ford F-350 XLT Pickup Truck Selling my 2015 Ford XLT F-350, Reg-Cab, 8' box, 45,000Km. I factory ordered this truck. The only add it doesn't have is the step tail gate, Ford warranty until 2020, also Ford maintenance package,… 45,435km | Automatic' added.
	// +: Simhash of 832df1ef4eb2e3e for ' 2016 Ford Mustang 2016 Ford Mustang white with black stripes, this car is in showroom shape and it only has 14,000kms. this beast has never been in an accident nor does it have one scratch on the body. i purchased 20… 14,000km | Automatic' added.
	// +: Simhash of 1772193fe56b4e4a for ' 2013 Ford Fusion Sedan 2013 Ford Fusion SE Ecoboost only 65KMs, comes equipped with black leather interior, rear view camera, navigation system, touch screen display, sunroof, heated front seats, dual climate control,… 65,665km | Automatic' added.
	// +: Simhash of 81cf91fe1eb4b2f for ' 2014 Ford Fiesta SE Hatchback 9500 OBO In need of money, selling my 2014 Ford Fiesta SE with 57000km on it, all electric, alloy wheels, very little gas consumption, very spacious, automatic transmission, front wheel drive FWD, bluetooth… 57,000km | Automatic' added.
	// +: Simhash of e8d6de1ee669f4fa for ' Trade 2016 Ford F-150 chrome bumpers for your black ones I'm looking to trade my chrome bumpers and grill of your black bumpers and grill. Please call or text 1,234km | Automatic' added.
	// =: Simhash of c833835fb8ef4733 for ' MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic' ignored.
	// =: Simhash of 58b2c31ce3ef3abd for ' ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic' ignored.
	// =: Simhash of c7ad53ee7a9354e for ' FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic' ignored.
	// +: Simhash of d8b2e11eebeb2df9 for ' 2016 Ford Focus SE Sedan This is the fabled old lady's car. Bought new Nov. 2016 and driven only 2,000 klms. before she stopped driving for good. This car is brand new and very clean. Still has that new car smell. Please… 2,035km | Automatic' added.
	// +: Simhash of c972bd1ce6e937fe for ' 2013 Ford F-150 XLT/XTR 4x4 3.5 Ecobust 2013 Ford F-150 XLT/XTR 4x4 3.5 Ecobust with 90000 km on it I'm a first owner. Always serviced at Dixie Ford Mississauga. Still on the warranty for 100k km or 5 years. Factory Max Tow package.… 90,000km | Automatic' added.
	// +: Simhash of 5830dd1fe4eb0f7e for ' 2012 Ford Fiesta Hatchback Selling my Ford fiesta, car runs in great conditions brand new breaks and front routers, excellent on consuming gas 58000 KM Smoke Free Safety is available contact 647 909 7578 Thanks 58,000km | Automatic' added.
	// +: Simhash of 4836637eb5eb223e for ' 2014 Ford Focus Automatic MINT CONDITION I am selling my 2014 focus. Excellent condition, very well maintained. Only 65000km. Economic both on Highway and city. Bluetooth, traction control, heated seats and much more. Call me or text me at… 65,000km | Automatic' added.
	// +: Simhash of 983ad71ea7eb37bc for ' 2014 Ford Focus SE Sedan Great car! No accidents! Low mileage! Winter tires in price! It runs perfectly, everything in a perfect condition - You should come and see it! This car is driven by my wife and need a bigger one. 25,734km | Automatic' added.
	// =: Simhash of 58b2c31ce3ef3abd for ' ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic' ignored.
	// +: Simhash of d8f8d73ea1e90c2c for ' 2014 Ford Mustang V6 Premium 2014 Ford Mustang with Pony package. 304 horsepower. leather power seat. Bluetooth. three drive model:comfort normal sport. 18 tires. two rims for you. the original rims are brand new in the box.… 41,500km | Automatic' added.
	// =: Simhash of d8f8d73ea1e90c2c for ' 2014 ford Mustang v6 premium 2014 Ford Mustang with Pony package. 304 horsepower. leather power seat. Bluetooth. three drive model:comfort normal sport. 18 tires. two rims for you. the original rims are brand new in the box.… 41,500km | Automatic' ignored.
	// +: Simhash of 983ab95cebe2376e for ' 2012 Ford Fiesta SES Ford Fiesta SES red leather interior fully loaded. Excellent condition with very low mileage. Never winter driven. Made for city driving, really good on gas. 9 months of financing at 0% interest… 35,100km | Automatic' added.
	// =: Simhash of d8f8d73ea1e90c2c for ' 2014 Ford Mustang Premium V6 2014 Ford Mustang with Pony package. 304 horsepower. leather power seat. Bluetooth. three drive model:comfort normal sport. 18 tires. two rims for you. the original rims are brand new in the box.… 41,500km | Automatic' ignored.
	// +: Simhash of c8f6fb3de9ef3f5e for ' 2014 Ford Fusion Titanium Sedan Very clean former salesman's car, mostly highway-driven. Safety certificate, all-season tires, just detailed. Fully-loaded with all the options Ford offered in 2014. Priced to move at $19,699 + HST 77,000km | Automatic' added.
	// +: Simhash of 883a935ebeeb094f for ' 12 Ford Focus SE, Low 58km Only! Clean history, no accident Winter tire & Summer tire Just got marry, need a bigger car for family!!! Text please, Jeremy Thanks! 58,000km | Automatic' added.
	// +: Simhash of 832b907a4eb15db for ' 2017 Ford Mustang GT CALIFORNIA SPECIAL Coupe (2 door) 2017 MUSTANG GT COUPE CALIFORNIA SPECIAL GRABBER BLUE INCLUDED ON THIS VEHICLE EQUIPMENT GROUP 401A SINGLE CD W/TUNER/SDARS/HD 12 SPEAKER SYSTEM OPTIONAL EQUIPMENT 5.0L 4V TI-VCT V8 ENGINE 6-SPEED… 5,540km | Automatic' added.
	// +: Simhash of c832173fa5eb283c for ' Low Mileage 2014 Ford Focus SE - Sport Package Hardly used with less than 15 thousand kilometers. Bought new in 2014 and in perfect condition. Heated seats and winter tires (on rims) included. Please call or text if interested. 14,500km | Automatic' added.
	// +: Simhash of cc3a491ef7eb3e20 for ' 2014 ford fusion Wrote off my car and this is the extra stuff in my garage that I need to geto rid of. Weather tech floor mats front+ back Factory rubber floor mats front+ back Factory carpet floor mats front+ back… 90,000km | Automatic' added.
	// +: Simhash of 9872e57ee1eb263d for ' 2013 Ford Escape SUV, Crossover VERY LOW KMs (37,500) FORD ESCAPE FOR SALE. This car has been very well maintained over its life and is 100% ACCIDENT FREE. Leather interior, heated seats = a comfortable ride. Only blemish on it is… 37,700km | Automatic | CarProof' added.
	// +: Simhash of c832d33fe7e91d43 for ' 2014 Ford Fusion SE Sedan-Amazing Car! Eco boost 2.0L I4 GTDI Engine Upgraded 18 Alloy Rims Set of Winter Tires Included (new 2015) Remote Keyless Entry New weather tech mats included Sync Voice Activated System Navigation System Rear… 51,500km | Automatic' added.
	// +: Simhash of c872fb1ee6a764ac for ' 2016 Ford F-250 Lariat Pickup Truck Diesel 2016 f250. Mint shape just no longer want payments. Looking for someone to take over. Fully loaded, has every option possible. 6.7 powerstroke, dpf and egr removed. Will come with dpf and egr. Has an… 40,000km | Automatic' added.
	// +: Simhash of 8b2df0ea6eb2f3c for ' 2015 Ford Explorer Sport SUV, Crossover This vehicle is a real beauty and a pleasure to drive. It is in excellent condition and has been store inside since purchased in 2015. It has not been driven in winter other then to go for service.!… 18,600km | Automatic' added.
	// +: Simhash of 8329706e4eb2f3d for ' 2013 Ford Fiesta Sedan - 22,116 kms Body is in perfect condition. No mechanical problems. Oil change and maintenance package done in March/17. Registered inspection done in April/16. $10,000 firm (sales tax is extra). Call … 22,120km | Automatic' added.
	// =: Simhash of c7ad53ee7a9354e for ' FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic' ignored.
	// =: Simhash of c833835fb8ef4733 for ' MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic' ignored.
	// =: Simhash of 58b2c31ce3ef3abd for ' ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic' ignored.
	// +: Simhash of 8d77f55ea66175ce for ' 2014 Ford E-250 Minivan, Van 2014 E250 Super Duty Regular Box 4.6L with A/C, bucket reclining seats, full cargo divider, side barn door, plywood floor, 3M protection and serviced regularly. Looking for some to assume lease or… 87,750km | Automatic' added.
	// +: Simhash of 9a72753ee0e10dd6 for ' 2013 Ford Focus SE Sedan 2013 December Ford Focus SE sedan for sale . Low kilometres... only 9000 kilometers. Looks like brand new. Good condition. Please contact me on 9,000km | Automatic' added.
	// +: Simhash of 32ff0fa4e33e3e for ' 2017 Ford Mustang EcoBoost Premium Coupe (2 door) I am looking to sell my 2017 Ford Mustang EcoBoost Premium. It is Grabber Blue in color. Driven by a professional accountant. Bought it in September 2016. Package details included: Premium package… 9,800km | Automatic' added.
	// +: Simhash of 834d55ee4eb1dae for ' 2013 Ford Fiesta SE Sedan Beautiful light blue Fiesta SE lady driven, non-smoking. Garage stored in the winter with no accidents. Great on gas and easy to park. Lots of options including LED lights for safe driving at night,… 87,000km | Automatic' added.
	// +: Simhash of 48d4ff3ea56b2f5f for ' 2012 Ford Mustang Premium Sport Coupe Selling my rare 2012 Club of America Edition Mustang. The car is fully loaded with all options available. Car is very well maintained, always stored in garage year round and comes with a separate set… 75,000km | Automatic' added.
	// =: Simhash of 48d5ff3ea56b2f5f for ' 2012 Ford Mustang Premium Sport Coupe (2 door) Selling my rare 2012 Club of America Edition Mustang. The car is fully loaded with all options available. Car is very well maintained, always stored in garage year round and comes with a separate set… 75,000km | Automatic' ignored.
	// +: Simhash of 836ff0ee6eb379e for ' 2010 Ford Mustang Hi I'm selling my 2010 ford mustang mint condition Power windows Power locks A/c 2 keys Back up sensors Set of tires on rims Safety and emission Car proof is provided 55,254km | Automatic' added.
	// +: Simhash of c83a771ea6eb167e for ' 2014 Ford Focus Hatchback Moving out of the country. Have to sell. Mint condition with lots of add-ons + new winter tires. SE Hatch with Sports package. only 27 months on the road. Includes maintenance package until March… 25,000km | Automatic' added.
	// +: Simhash of d83cf77ee3f897ec for ' 2016 Ford Mustang Convertible 2016 Mustang in beautiful shape clean and like new. Never winter driven. One owner-bought brand new. No accidents and tires are like new with 5000 kilometers. electric seat, mirrors, keyless,… 5,000km | Automatic' added.
	// +: Simhash of a73e75ee7e917ef for ' Ford Focus 2011 SE for sale Blue 4 Door Sedan, mint condition, low mileage (under 50,000km) 49,000km | Automatic' added.
	// +: Simhash of d89ec51ee3ebb6a8 for ' Ford F-150 EcoBoost 2014 crew cab with box liner and cover. Low kilometers. Back up camera and sensors. Satellite radio and CD player cold AC and rubber floor mats. Cloth seats. Six seater, remaining bumper to bumper… 26,000km | Automatic' added.
	// +: Simhash of 187bb73ee1e0967e for ' Lease takeover!! 2016 Ford Edge- fwd black on black leather Automatic keyless entry, automatic trunk, panoramic sunroof, push start, automatic start with key. 24 months left on lease with 14,300 kms. I am currently… 14,300km | Automatic' added.
	// +: Simhash of d8fec71eeba93f7e for ' 2012 Ford Focus SEL Hatchback for Sale - 1 owner Ford Focus Hatchback for sale by owner. Gently used with low mileage. Maintenance has been done up to date. Recent oil change and disc brake replacement. The car has been detailed and cleaned for the… 76,261km | Automatic' added.
	// +: Simhash of 98727b1fe6e8244f for ' 2014 Ford F-150 XL Pickup Truck sold by owner 2014 Ford F-150 XL Pickup Truck sold by owner, low mileage (64000km). New tires, cap, hitch, no accidents, great conditions. First come first buy. Please, contact me 64,000km | Automatic | CarProof' added.
	// +: Simhash of d8f8e03ea7ef177e for ' Ford Mustang Very clean vehicle comes with summer and winters on rims all dealer serviced paper work and all rust proofed every year breaks are brand new to start off the summer warranty still available I have a… 55,000km | Automatic | CarProof' added.
	// +: Simhash of 872f07ef6a966ac for ' 2012 Ford F-250 Super Duty XL Crew Cab Pickup Truck Commercial Truck Like New 2012 Ford F250 Super Duty XL Crew Cab 6.5' Box Grey 80K. Truck truly like Brand New & Very Clean and Powerful, it's barely used and had no accident whatsoever, also you'll… 80,000km | Automatic' added.
	// +: Simhash of 873fd2ee46b5c0f for ' 2015 Ford F-150 Lariat SuperCrew Pickup Truck 2015 Ford F150 Lariat. Loaded. No sunroof. 500A Available August 1st, 2017 Presently in northern Ontario. Contact Martin 55,000km | Automatic' added.
	// =: Simhash of 873fd2ee46b5c0f for ' 2015 Ford F-150 Lariat SuperCrew Pickup Truck 2015 Ford F150 Lariat. Loaded. No sunroof. 500A Available August 1st, 2017 Presently in northern Ontario. Contact Martin 55,000km | Automatic' ignored.
	// +: Simhash of 4df3bf0ea7c925dc for ' 2015 Ford F-150 SuperCrew FX4 Off Road sport Pickup Truck Mint condition 3.5 Eco Boost V6. Aluminum body. Still under warranty. Heated power Lombard seats. Brake pedal adjustment. Power windows, locks, mirrors, sunroof. Tilt and telescopic steering wheel.… 21,000km | Automatic' added.
	// +: Simhash of 4872dd1ee0eb4ab4 for ' Ford Fusion 2011 SE BLACK **CLEAN** 2011 Model Clean Ford Fusion SE Oil is newly changed. NO CRASH, ALL PARTS ORIGINAL. Only Left Mirror is Broken. (Still, works and can easily be replaced) AROUND 199.297 km If you want a test drive,… 199km | Automatic' added.
	// =: Simhash of c7ad53ee7a9354e for ' FORD 550 DUMP TRUCK AND BINS FOR SALE VEHICLE OPTIONS: Air Conditioning CD player Leather seats Airbag: passenger Tow package FORD F550 dump truck for sale also 9 bins 14 yard. 4 bins 18yard. Truck 2016 still got 5 year finance Monthly… 85,000km | Automatic' ignored.
	// =: Simhash of c833835fb8ef4733 for ' MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic' ignored.
	// =: Simhash of 58b2c31ce3ef3abd for ' ***2011 LINCOLN MKX AWD - ORIGINAL OWNER*** Selling our family vehicle which has been maintained from day one. We are the original owners, bought from East Court Ford. Condition is excellent, KMS are low at 80,000 and all service records to… 80,000km | Automatic' ignored.
	// +: Simhash of 9963bd0ce6013fcd for ' 2016 Ford F-150 Lariat Pickup Truck 2.7L 501A w/FX4 Offroad Pky Vehicle Features: 2016 Magnetic Grey Ford F-150 SuperCrew 145 Wheelbase FX4 Off Road Package w/Skid Plates, 3.73 Ratio E-lock 2.7L EcoBoost Twin-Turbo Engine 6-Speed SelectShift Transmission 6500#… 19,999km | Automatic' added.
	// +: Simhash of 98b6bc7ce3db0466 for ' 2015 Ford Escape Titanium SUV, Crossover Next to new Ford Escape Titanium model. Sun roof, 4WD, full set of winter tires & rims, power liftgate, Bluetooth, heated mirrors, steering wheel audio controls, rear window defogger, rearview camera… 41,193km | Automatic' added.
	// +: Simhash of 9c72bd1ea4eb367d for ' 2013 Ford F-150 SuperCrew Pickup Truck Great Clean truck, 301A package, Brakes just done, power mirrors, Pedals, Windows, Seats Bluetooth and satellite radio Trailer brake Controller and Class III hitch 80000 km 6.5 foot box 5.0L V8 motor… 80,000km | Automatic' added.
	// +: Simhash of 43726f1ee7eb010b for ' 2015 Ford Focus Sedan ***Reduced***2015 Ford Focus fully loaded, leather interior, power sun roof, full warranty, full maintenance package included, back up camera, list goes on. Fuel is amazing. Will sell or you may take… 55,000km | Automatic' added.
	// +: Simhash of 2176b94fe44b974c for ' Ford Mustang 2015 Ecoboost - Excellent Condition Ford Mustang Ecoboost - Excellent Condition 2015 Ford Mustang EcoBoost 2.3 L Automatic, 305 Horse Power, Absolute Beauty Keyless Entry, Remote Starter, Bluetooth, Steering Tilt, Reverse Camera, Power… 51,000km | Automatic' added.
	// +: Simhash of 832d50ef4eb3767 for ' 2013 Ford Escape SE Price is firm. Immaculate - no pets, no kids, non-smoker, female driven, I do not eat or drink in the car. Just got oil changed. I absolutely love this vehicle, just need to downsize to a car. 1.6L… 68,000km | Automatic' added.
	// =: Simhash of c833835fb8ef4733 for ' MINI EXCAVATOR SERVICES $90/HOUR 24 hour EMERGENCY SERVICES Mini excavator & skidsteer service with operator from $90/ hour. Fully insured and extremely experienced operator. EXCAVATING, GRADING TRENCHING AUGER CONCRETE BREAKER PLATE TAMPER SKIDSTEER SERVICE… 18,000km | Automatic' ignored.
	// +: Simhash of 9870fd3ca6e9371c for ' 2013 Ford F-450 King Ranch Pickup Truck FORD EXTENDED WARRANTY PREMIUM CARE PLAN - AIRLIFT REAR AIR SUSPENSION WITH IN CAB AIR CONTROLLER. THIS IS NIT A CHEAP AFTERMARKET PRODUCT.… 30,500km | Automatic' added.
	// +: Simhash of d8fec71aefe90e1f for ' XLT Beautiful SUV. Has XLT package. Very low Kms's on it and hardly used. Message for details. 25,000km | Automatic' added.
	// +: Simhash of 872657fe2eb1d2f for ' 2013 Ford Escape SE SUV, Crossover Selling a 2013 Ford Escape SE fully loaded still under warranty from Ford Canada. Need to sell the car ASAP as I have no place to keep at home. Contact: pagal_boi@hotmail.com 80,000km | Automatic' added.
	// +: Simhash of 872531eb6e9040e for ' -- URGENT-- 2010 Ford F-150 XL -- MINT -- BULLET PROOF V8 — Hello, I have my beautiful Grey Ford F150 XL for sale. This truck has been driven by a senior to the grocery store and back. This truck has never been smoked in and is very clean base model. It has a… 41,500km | Automatic' added.
	// +: Simhash of 8c73fd7c86291740 for ' Ford Escape Lease take over, Ford Escape 2017,1.5 L /ecoboost 41 months left, payment is 420CAN per month. Very clean like a new car 8,000km | Automatic' added.
	// +: Simhash of 36757fe56b1c29 for ' 2014 Ford Escape SE SUV, Crossover 4X4, NAVIGATION, BACK UP CAMERA, BLUETOOTH, ROOF RAIL, SUNROOF. Rebuilders title. CALL 40,200km | Automatic' added.
	// +: Simhash of 8832d71eb4eb3fad for ' 2010 Ford Focus SE LOW KM 2010 Ford Focus SE with only 79,000 KM Great little car, no longer need car as not commuting to school, lady driven, Second Owner. Features: (listed as many as I could think of, email any questions… 79,000km | Automatic' added.
	// +: Simhash of c036cf2de5fb2e2c for ' 2014 Ford Escape Titanium SUV, Crossover 4WD TITANIUM, POWER HEATED LEATHER SEATS, PANORAMIC SUNROOF, PREMIUM AUDIO WITH BLUETOOTH PHONE CONNECTIVITY, REMOTE STARTER, BACK-UP CAMERA AND MUCH MORE. THIS PREMIUM ECOBOOST IS IN EXCELLENT SHAPE… 54,400km | Automatic' added.
	// +: Simhash of 8832dd0f24eb0049 for ' 2014 Ford Escape SE SUV, Lease takeover $344.00 Lease takeover. 11 months remaining on lease. Return date : May 28, 2018 80,000 km lease allowance. Current km reading 49,500 $344.00 per month - plus tax Pearl White, SE, 4wd, panoramic sunroof,… 49,500km | Automatic' added.
	// +: Simhash of 88327f26f4eb0c11 for ' Mint 2014 Ford F-250 New tires Motor is good Coolant is leaking Transmission is good Selling as is 330km | Automatic' added.
	// +: Simhash of 59f7ee7fe4ffa544 for ' 2015 Ford Escape 4WD 4dr SE SUV, Crossover Dual front climate control, Power windows, Power door locks, Power drivers seat, Power mirrors, Cloth seats, Heated seats, Bucket seats, Split bench seat, 17 inch alloy wheels, All season tires,… 23,000km | Automatic' added.
	// +: Simhash of 8833db1ea6e9371e for ' Lease takeover - 2017 Mustang GT Premium w/ California Special Looking for someone to take over my lease, I am moving to the States and I cannot take this vehicle with me. Car was leased Oct/16 with 25000km annual milage. Contact me for more information. 8,950km | Automatic' added.
	// +: Simhash of 4850d27deaad0fa4 for ' 2011 Ford Escape Hatchback Clean and reliable vehicle, never had any problems with it. Comes with winter tires 85,000km | Automatic' added.
	// ================
	// +: Simhash of c833835fb8ef4733 added.
	// +: Simhash of c7ad53ee7a9354e added.
	// +: Simhash of 58b2c31ce3ef3abd added.
	// +: Simhash of 5876d51ea5eb2f7c added.
	// +: Simhash of d8fafe1ee3eb3e0d added.
	// +: Simhash of 5976ee1ea6ef375e added.
	// +: Simhash of 8873df0fe6eb36bc added.
	// +: Simhash of 1832c51ee6eb2e3e added.
	// +: Simhash of 872f71fe0ebba3a added.
	// +: Simhash of d8dcc1186ba977be added.
	// +: Simhash of 8d63bf7ee661affc added.
	// +: Simhash of d89abd1ceee92e7c added.
	// +: Simhash of 9872bf7ee2e13425 added.
	// +: Simhash of 4472bd7fe4eb154c added.
	// +: Simhash of 4872e50fa6eb3576 added.
	// +: Simhash of 9932b71ea4a934bd added.
	// +: Simhash of c1b2ef3fe1b95f0e added.
	// +: Simhash of 8f2f34ea4eb364f added.
	// +: Simhash of 897af11ee6e9146f added.
	// =: Simhash of 832df1ef4eb2e3e ignored for 1832c51ee6eb2e3e (6).
	// +: Simhash of 1772193fe56b4e4a added.
	// +: Simhash of 81cf91fe1eb4b2f added.
	// +: Simhash of e8d6de1ee669f4fa added.
	// =: Simhash of c833835fb8ef4733 ignored for c833835fb8ef4733 (0).
	// =: Simhash of 58b2c31ce3ef3abd ignored for 58b2c31ce3ef3abd (0).
	// =: Simhash of c7ad53ee7a9354e ignored for c7ad53ee7a9354e (0).
	// +: Simhash of d8b2e11eebeb2df9 added.
	// +: Simhash of c972bd1ce6e937fe added.
	// +: Simhash of 5830dd1fe4eb0f7e added.
	// +: Simhash of 4836637eb5eb223e added.
	// +: Simhash of 983ad71ea7eb37bc added.
	// =: Simhash of 58b2c31ce3ef3abd ignored for 58b2c31ce3ef3abd (0).
	// +: Simhash of d8f8d73ea1e90c2c added.
	// =: Simhash of d8f8d73ea1e90c2c ignored for d8f8d73ea1e90c2c (0).
	// +: Simhash of 983ab95cebe2376e added.
	// =: Simhash of d8f8d73ea1e90c2c ignored for d8f8d73ea1e90c2c (0).
	// +: Simhash of c8f6fb3de9ef3f5e added.
	// +: Simhash of 883a935ebeeb094f added.
	// +: Simhash of 832b907a4eb15db added.
	// +: Simhash of c832173fa5eb283c added.
	// +: Simhash of cc3a491ef7eb3e20 added.
	// +: Simhash of 9872e57ee1eb263d added.
	// +: Simhash of c832d33fe7e91d43 added.
	// +: Simhash of c872fb1ee6a764ac added.
	// +: Simhash of 8b2df0ea6eb2f3c added.
	// =: Simhash of 8329706e4eb2f3d ignored for 8b2df0ea6eb2f3c (7).
	// =: Simhash of c7ad53ee7a9354e ignored for c7ad53ee7a9354e (0).
	// =: Simhash of c833835fb8ef4733 ignored for c833835fb8ef4733 (0).
	// =: Simhash of 58b2c31ce3ef3abd ignored for 58b2c31ce3ef3abd (0).
	// +: Simhash of 8d77f55ea66175ce added.
	// +: Simhash of 9a72753ee0e10dd6 added.
	// +: Simhash of 32ff0fa4e33e3e added.
	// +: Simhash of 834d55ee4eb1dae added.
	// +: Simhash of 48d4ff3ea56b2f5f added.
	// =: Simhash of 48d5ff3ea56b2f5f ignored for 48d4ff3ea56b2f5f (1).
	// +: Simhash of 836ff0ee6eb379e added.
	// +: Simhash of c83a771ea6eb167e added.
	// +: Simhash of d83cf77ee3f897ec added.
	// +: Simhash of a73e75ee7e917ef added.
	// +: Simhash of d89ec51ee3ebb6a8 added.
	// +: Simhash of 187bb73ee1e0967e added.
	// +: Simhash of d8fec71eeba93f7e added.
	// +: Simhash of 98727b1fe6e8244f added.
	// +: Simhash of d8f8e03ea7ef177e added.
	// +: Simhash of 872f07ef6a966ac added.
	// +: Simhash of 873fd2ee46b5c0f added.
	// =: Simhash of 873fd2ee46b5c0f ignored for 873fd2ee46b5c0f (0).
	// +: Simhash of 4df3bf0ea7c925dc added.
	// +: Simhash of 4872dd1ee0eb4ab4 added.
	// =: Simhash of c7ad53ee7a9354e ignored for c7ad53ee7a9354e (0).
	// =: Simhash of c833835fb8ef4733 ignored for c833835fb8ef4733 (0).
	// =: Simhash of 58b2c31ce3ef3abd ignored for 58b2c31ce3ef3abd (0).
	// +: Simhash of 9963bd0ce6013fcd added.
	// +: Simhash of 98b6bc7ce3db0466 added.
	// +: Simhash of 9c72bd1ea4eb367d added.
	// +: Simhash of 43726f1ee7eb010b added.
	// +: Simhash of 2176b94fe44b974c added.
	// +: Simhash of 832d50ef4eb3767 added.
	// =: Simhash of c833835fb8ef4733 ignored for c833835fb8ef4733 (0).
	// +: Simhash of 9870fd3ca6e9371c added.
	// +: Simhash of d8fec71aefe90e1f added.
	// +: Simhash of 872657fe2eb1d2f added.
	// +: Simhash of 872531eb6e9040e added.
	// +: Simhash of 8c73fd7c86291740 added.
	// +: Simhash of 36757fe56b1c29 added.
	// +: Simhash of 8832d71eb4eb3fad added.
	// +: Simhash of c036cf2de5fb2e2c added.
	// +: Simhash of 8832dd0f24eb0049 added.
	// +: Simhash of 88327f26f4eb0c11 added.
	// +: Simhash of 59f7ee7fe4ffa544 added.
	// +: Simhash of 8833db1ea6e9371e added.
	// +: Simhash of 4850d27deaad0fa4 added.
	// ================
}
