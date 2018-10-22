package main

import (
	"github.com/goGZipJSONVerification/common"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8091", common.HandlerCompressor(handlerPersonal)))
}

func handlerPersonal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Accept-Charset", "utf-8")
	w.Write([]byte(getJsonData()))
}

func getJsonData() string {
	return `{
	"id": 56921682,
	"address": {
		"id": 2089195,
		"area": {
			"id": 18536,
			"city": {
				"id": 123,
				"name": "Belo Horizonte",
				"state": {
					"id": 35,
					"name": "MG"
				}
			},
			"name": "São Salvador"
		},
		"complement": "Casa verde dois pavimentos",
		"corner": "Bar da Viúva, Salão Varanda",
		"doorNumber": "129",
		"notes": "",
		"phone": "998808690",
		"street": "Líbia",
		"zipCode": ""
	},
	"addressCoordinates": "-19.9007,-44.0148",
	"addressDescription": "Líbia 129 Casa verde dois pavimentos esquina Bar da Viúva, Salão Varanda",
	"addressNotes": "",
	"addressPhone": "998808690",
	"addressZipCode": "",
	"amountNoDiscount": 31.99,
	"application": "ANDROID",
	"clientGuid": "db6aa198-924f-4c12-a80c-bcfe31c43272",
	"code": "69OAHOH7",
	"commission": 12,
	"commissionIncludesDeliveryCost": true,
	"creditCardCommission": 0,
	"data": {
		"id": 57010175,
		"first": false,
		"ipAddress": "2804:14c:5bd2:8c85:7dbd:e42b:efa2:b245",
		"isExpress": false,
		"parallelReceptionSystem": false,
		"receptionSystem": {
			"id": 57,
			"code": "desktop_app",
			"name": "Desktop Application",
			"push": false
		},
		"receptionSystemEnabled": true,
		"secondaryReceptionSystem": {
			"id": 2,
			"code": "call_center",
			"name": "Call center",
			"push": false
		}
	},
	"deliveryDate": "2018-10-14T18:51:06Z",
	"deliveryZoneId": 312558,
	"details": [
		{
			"id": 111020775,
			"discountAmount": 0,
			"optionGroups": [
				{
					"id": 100692999,
					"name": "Turbine seu burguer",
					"productOptionGroup": {
						"id": 1889233,
						"index": 0
					}
				}
			],
			"product": {
				"id": 2656922,
				"foodCategoryTag": {
					"id": 2371,
					"foodCategory": {
						"id": 115,
						"name": "Hambúrguer"
					},
					"name": "Hamburguer"
				},
				"name": "05 - Hambúrguer dança com lobos",
				"price": 24.99,
				"section": {
					"id": 318455,
					"index": 1,
					"name": "Hamburguers"
				}
			},
			"productName": "05 - Hambúrguer dança com lobos",
			"quantity": 1,
			"subtotal": 24.99,
			"taxAmount": 0,
			"taxRate": 0,
			"total": 24.99,
			"unitPrice": 24.99
		},
		{
			"id": 111020776,
			"discountAmount": 0,
			"optionGroups": [],
			"product": {
				"id": 2966010,
				"name": "Mate Couro 1 L",
				"price": 7,
				"section": {
					"id": 318459,
					"index": 5,
					"name": "Bebidas"
				}
			},
			"productName": "Mate Couro 1 L",
			"quantity": 1,
			"subtotal": 7,
			"taxAmount": 0,
			"taxRate": 0,
			"total": 7,
			"unitPrice": 7
		}
	],
	"discount": 0,
	"discountType": "NONE",
	"discounts": [],
	"foodTaxesAmount": 0,
	"hash": "6bf35f634ce948a8b026a7c53af691ab",
	"logisticsCommission": 0,
	"paymentAmount": 41.99,
	"paymentMethod": {
		"id": 37,
		"descriptionEN": "Ticket Restaurant",
		"descriptionES": "Ticket Restaurant",
		"descriptionPT": "Ticket Restaurante",
		"image": "payments/pay-ticket-restaurant.png",
		"name": "Ticket Restaurant",
		"online": false,
		"uiType": "TICKET"
	},
	"paymentNotes": "",
	"pickup": false,
	"promisedDeliveryTime": {
		"id": 3,
		"description": "Entre 45' y 60'",
		"maxMinutes": 60,
		"minMinutes": 45,
		"name": "Entre45Y60",
		"order": 3
	},
	"pushed": false,
	"registeredDate": "2018-10-14T18:51:06Z",
	"restaurant": {
		"id": 43766,
		"address": {
			"id": 4614073,
			"area": {
				"id": 1207,
				"city": {
					"id": 131,
					"name": "Contagem",
					"state": {
						"id": 35,
						"name": "MG"
					}
				},
				"name": "Cabral"
			},
			"doorNumber": "393",
			"latitude": -19.880599975585938,
			"longitude": -44.04180145263672,
			"phone": "5531-3394-7129",
			"street": "Alameda dos Sabiás"
		},
		"automaticPhone": "+553133947129",
		"automaticPhoneEnabled": true,
		"billingInfo": {
			"id": 39776,
			"companyNumber": "24897640000104"
		},
		"businessType": "RESTAURANT",
		"country": {
			"id": 5,
			"culture": "pt_BR",
			"currency": {
				"id": 5,
				"isoCode": "BRL",
				"symbol": "R$"
			},
			"name": "Brasil",
			"shortName": "br",
			"timeOffset": -180,
			"url": "https://www.pedidosja.com.br"
		},
		"deliveryTime": {
			"id": 3,
			"maxMinutes": 60,
			"minMinutes": 45
		},
		"foodCategories": [
			{
				"id": 12789489,
				"foodCategory": {
					"id": 120,
					"name": "Peixes e Frutos do Mar"
				},
				"percentage": 0.02
			},
			{
				"id": 12789494,
				"foodCategory": {
					"id": 115,
					"name": "Hambúrguer"
				},
				"percentage": 0.27
			},
			{
				"id": 12789492,
				"foodCategory": {
					"id": 97,
					"name": "Carnes"
				},
				"percentage": 0.05
			},
			{
				"id": 12789495,
				"foodCategory": {
					"id": 116,
					"name": "Lanches"
				},
				"percentage": 0.02
			},
			{
				"id": 12789488,
				"foodCategory": {
					"id": 110,
					"name": "Doces"
				},
				"percentage": 0.02
			},
			{
				"id": 12789487,
				"foodCategory": {
					"id": 123,
					"name": "Salgados"
				},
				"percentage": 0.08
			},
			{
				"id": 12789493,
				"foodCategory": {
					"id": 131,
					"name": "Espetinho"
				},
				"percentage": 0.02
			},
			{
				"id": 12789491,
				"foodCategory": {
					"id": 127,
					"name": "Sucos"
				},
				"percentage": 0.03
			},
			{
				"id": 12789496,
				"foodCategory": {
					"id": 95,
					"name": "Bebidas"
				},
				"percentage": 0.2
			},
			{
				"id": 12789490,
				"foodCategory": {
					"id": 119,
					"name": "Pastel"
				},
				"percentage": 0.02
			}
		],
		"hash": "84850a35b778487cba357a38e99d58b1",
		"isImportantAccount": false,
		"link": "django-hamburgueria",
		"logo": "django-hamburgueria.jpg",
		"name": "Django Hamburgueria",
		"ordersReceptionSystem": {
			"id": 57,
			"code": "desktop_app",
			"itWorks": true,
			"name": "Desktop Application",
			"push": false
		},
		"ordersSecondaryReceptionSystem": {
			"id": 2,
			"code": "call_center",
			"name": "Call center",
			"push": false
		},
		"parallelReceptionSystem": false,
		"publicPhone": "33947129",
		"ratings": [
			{
				"id": 6984,
				"generalScore": 4.32,
				"platform": "PEDIDOS_YA"
			}
		],
		"receptionSystemEnabled": true,
		"registeredDate": "2017-03-24T21:24:01Z"
	},
	"shippingAmount": 10,
	"shippingAmountNoDiscount": 10,
	"shippingTaxesAmount": 0,
	"stampsDiscount": 0,
	"state": "PENDING",
	"taxAmount": 0,
	"totalAmount": 31.99,
	"user": {
		"id": 31,
		"avatar": "https://graph.facebook.com/611635355641497/picture?width=100&height=100",
		"country": {
			"id": 5,
			"name": "Brasil"
		},
		"data": {
			"id": 1878396,
			"birth": "2000-08-05T00:00:00Z",
			"gender": "MALE",
			"orderCount": 342
		},
		"email": "gabrieluiz2000@gmail.com",
		"firstOrderDate": "2015-12-05T10:53:52Z",
		"hash": "eb704dd437d348dbb61fb34166087e39",
		"identityCard": "53092023668",
		"lastName": "Souza",
		"mobile": "",
		"name": "Gabriel",
		"platform": "PEDIDOS_YA",
		"receivesPushReviewReminder": true,
		"receivesReviewReminder": true,
		"receivesWhatsappNotifications": true,
		"registeredDate": "2015-12-05T10:42:47Z",
		"type": "WEB_USER",
		"userTrustScore": {
			"id": 2118723,
			"confirmedOfflineOrderCount": 342,
			"confirmedOnlineOrderCount": 0
		},
		"whatsAppUserId": "553198808690",
		"facebookId": "611635355641497"
	},
	"userIdentityCard": "53092023668",
	"withLogistics": false,
	"deliveryExpectedFrom": "2018-10-14T19:36:06Z",
	"deliveryExpectedTo": "2018-10-14T19:51:06Z",
	"deviceToken": "dcnkWfoD89Q:APA91bGnhZLWeHYVQC8B2Dh-C8AYIzX6O5p_Rm4L8Pztj0cSGQwjChwmaQXIJdGBqOlpJmw-fsw4XO-JZCszoNERTWjWb6VNzm6wBnNE4wI7I1Clfa3s6jty_auSM-Dc9nliCc2LILXxptbYF9XMzskho5MYohIQkg"
}`
}
