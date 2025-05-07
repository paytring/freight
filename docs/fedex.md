# auth


`grant_type=client_credentials&client_id=l75c4815a4147d44e6af1de680e269508c&client_secret=e5190e31828b41409d14cebc4bedcba8`
## request
```bash
curl --location 'https://api.fedex.com/auth/oauth/v2/token' \
--header 'accept: application/json, text/plain, */*' \
--header 'accept-language: en-GB,en-US;q=0.9,en;q=0.8' \
--header 'cache-control: no-cache' \
--header 'content-type: application/x-www-form-urlencoded;charset=UTF-8' \
--header 'origin: https://www.fedex.com' \
--header 'referer: https://www.fedex.com/secure-login/en-in/' \
--data 'grant_type=client_credentials&client_id=l75c4815a4147d44e6af1de680e269508c&client_secret=e5190e31828b41409d14cebc4bedcba8'

```

## response
```json
{
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJDWFMtVFAiXSwiUGF5bG9hZCI6eyJjbGllbnRJZGVudGl0eSI6eyJjbGllbnRLZXkiOiJsN2JjN2QxM2IwOTQ2NjRhZDlhMDhlMDFjZTlhMjBmZTRjIn0sImF1dGhlbnRpY2F0aW9uUmVhbG0iOiJDTUFDIiwiYWRkaXRpb25hbElkZW50aXR5Ijp7InRpbWVTdGFtcCI6IjA1LU1heS0yMDI1IDAzOjIzOjU5IEVTVCIsImdyYW50X3R5cGUiOiJjbGllbnRfY3JlZGVudGlhbHMiLCJhcGltb2RlIjoiU2FuZGJveCIsImN4c0lzcyI6Imh0dHBzOi8vY3hzYXV0aHNlcnZlci1zdGFnaW5nLmFwcC5wYWFzLmZlZGV4LmNvbS90b2tlbi9vYXV0aDIifSwicGVyc29uYVR5cGUiOiJEaXJlY3RJbnRlZ3JhdG9yX0IyQiJ9LCJleHAiOjE3NDY0MzcwMzksImp0aSI6ImFmNGY2OTE4LTIzZTQtNDFjOC04OTllLWYxMWMxNmJmYjg3MCJ9.wWZoGxWbOJx1rgzi0lMzYZZtCikpHrbMkJQGP4ZRhcxYOAoYWcjafT0sOrWTkoPyiEzEBeeLv1w0P2jg0lt5DLADT9rSSMAoqM_Ss1Fo_fHcQKc8yFOeyOw2lSJfQqV3_b7O2U1yYGJeFkB3it2Hqu0WbKW6DtfoDEguMkPStYw0ydxangdmjiGUyxKXGR6LqgIWqERM1TjlQPcZG6KnDNwrzf2mBj5DpefUuCMUXQm9LYGTY2FKvqhjvzhzh9HvDAcniBDpH69Ay188TxW9o8PhriGldlOmQizXwJvVb9ZYh0C58g-y2LCg7-OxdHllUq4RPYj4WCRdTQ_-Ty-ov-g3l6tCIY9Y0-bZipdG2_XZWtcG8QOutLTQI5MgcVZeD0FRbnnKX00GRkeCrQtFqwG8TWMJxXJ0B9pxVDq0EQ-v3I77Z1-GYDAb2tlbZ_-Scrv2U9n5MfaeHybp8uK4Cx2qPkKV5XLLqxIQNXNtfmI8a62KKdWzIG7eV3dGbppXtO48wfkvgfheVZGCVSPyltef9Scqm6IUgL4V6el9Vnao-Z3hY52ChCBECBEeviH1yJu6_-_t1R6DbOdtaHymT8SYhyrXlEJr7D6UQEKmic6PwvGfS9fOFYS-g_BWoU5oBtCxZEDxCMfUQlsDETypP2qEu4dIp6k9g-6-Eh_cZHk",
    "token_type": "bearer",
    "expires_in": 3599,
    "scope": "CXS-TP"
}
```

# rates api

## request

```bash
curl --location 'https://api.fedex.com/rate/v2/rates/quotes' \
--header 'authorization: Bearer l7bfc00109fa7c4522bca5b1bc6fd6bf0d' \
--header 'origin: https://www.fedex.com' \
--header 'referer: https://www.fedex.com/en-in/online/rating.html' \
--header 'x-clientid: MAGR' \
--header 'x-clientversion: 1.0' \
--header 'x-locale: en_IN' \
--header 'Content-Type: application/json' \
--data '{
    "rateRequestControlParameters": {
        "rateSortOrder": "COMMITASCENDING",
        "returnTransitTimes": true,
        "variableOptions": null,
        "servicesNeededOnRateFailure": false
    },
    "requestedShipment": {
        "shipper": {
            "accountNumber": {
                "key": "a07ef98b62be7bca88e2007220c23084",
                "value": "453673782"
            },
            "address": {
                "city": "Chennai",
                "postalCode": "600003",
                "countryCode": "IN",
                "residential": false,
                "stateOrProvinceCode": "TN"
            }
        },
        "recipients": [
            {
                "address": {
                    "city": "Oklahoma City",
                    "postalCode": "73102",
                    "countryCode": "US",
                    "residential": false,
                    "stateOrProvinceCode": "OK"
                }
            }
        ],
        "shipTimestamp": "2025-05-08",
        "pickupType": "DROPOFF_AT_FEDEX_LOCATION",
        "packagingType": "FEDEX_10KG_BOX",
        "shippingChargesPayment": {
            "payor": {
                "responsibleParty": {
                    "accountNumber": {
                        "key": "a07ef98b62be7bca88e2007220c23084",
                        "value": "453673782"
                    },
                    "address": {
                        "countryCode": "IN"
                    }
                }
            }
        },
        "blockInsightVisibility": false,
        "edtRequestType": "NONE",
        "rateRequestType": [
            "ACCOUNT",
            "LIST"
        ],
        "requestedPackageLineItems": [
            {
                "groupPackageCount": 1,
                "physicalPackaging": "FEDEX_10KG_BOX",
                "insuredValue": {
                    "currency": "INR",
                    "currencySymbol": null,
                    "amount": 0
                },
                "weight": {
                    "units": "KG",
                    "value": 0.5
                }
            }
        ],
        "preferredCurrency": "INR",
        "customsClearanceDetail": {
            "commodities": [
                {
                    "name": "NON_DOCUMENTS",
                    "numberOfPieces": 1,
                    "weight": {
                        "units": "KG",
                        "value": 0.5
                    },
                    "quantity": 1,
                    "quantityUnits": "",
                    "unitPrice": {
                        "currency": "INR",
                        "amount": null,
                        "currencySymbol": ""
                    },
                    "customsValue": {
                        "currency": "INR",
                        "amount": "1.000",
                        "currencySymbol": ""
                    }
                }
            ]
        }
    },
    "carrierCodes": [
        "FDXG",
        "FDXE"
    ],
    "returnLocalizedDateTime": true,
    "webSiteCountryCode": "IN"
}''
```

## response
```json
{
    "transactionId": "55548914-8aec-4daf-bd17-4701458ad883",
    "output": {
        "rateReplyDetails": [
            {
                "serviceType": "FEDEX_INTERNATIONAL_PRIORITY_EXPRESS",
                "serviceSubOptionDetail": {},
                "serviceName": "FedEx International Priority® Express",
                "packagingType": "FEDEX_10KG_BOX",
                "commit": {
                    "dateDetail": {
                        "dayOfWeek": "Sat",
                        "time": "13:30",
                        "day": "May 10, 2025",
                        "dayCxsFormat": "May-10-2025"
                    },
                    "commodityName": "BOOKS"
                },
                "commitDetails": [
                    {
                        "dateDetail": {
                            "dayOfWeek": "Sat",
                            "time": "13:30",
                            "day": "May 10, 2025",
                            "dayCxsFormat": "May-10-2025"
                        },
                        "commodityName": "BOOKS"
                    }
                ],
                "customerMessages": [
                    {
                        "code": "customsnote",
                        "message": "<b>Important for Customs:</b> To prevent delays, your shipment must have the following customs documents attached: Commercial Invoice, Certificate of Origin, Shippers Export Declaration. Additional clearance documents may also be required. For more information, please visit our <A HREF=\"https://www.fedex.com/GTM?cntry_code=IN/\">international tools</A> site."
                    },
                    {
                        "code": "postalAwareAdditionalNote",
                        "message": "For countries where zip/postal code or city name can be entered, it is recommended that zip/postal code be entered to obtain a more accurate number of available services."
                    },
                    {
                        "code": "multiweightNote",
                        "message": "Multiweight rated shipment"
                    }
                ],
                "ratedShipmentDetails": [
                    {
                        "rateType": "ACCOUNT",
                        "ratedWeightMethod": "ACTUAL",
                        "effectiveNetDiscount": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 16339.7
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 27281.9
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 23120.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 27281.9
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2A_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 27281.9
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 2080.8
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 2080.8
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "SATURDAY_DELIVERY",
                                        "description": "Saturday Delivery",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1500.0
                                            }
                                        ]
                                    },
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 5197.6
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 6780.6
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "SATURDAY_DELIVERY",
                                    "description": "Saturday Delivery",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 1500.0
                                        }
                                    ]
                                },
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 5197.6
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "effectiveNetDiscount": {
                                    "currency": "INR",
                                    "amount": 0.0
                                },
                                "packageRateDetail": {
                                    "rateType": "PAYOR_ACCOUNT_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 6780.6
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 23120.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 4161.6
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 27281.9
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "SATURDAY_DELIVERY",
                                            "description": "Saturday Delivery",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 1500.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 5197.6
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    },
                    {
                        "rateType": "LIST",
                        "ratedWeightMethod": "ACTUAL",
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 16339.7
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 27281.9
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 23120.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 27281.9
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2A_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 27281.9
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 2080.8
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 2080.8
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "SATURDAY_DELIVERY",
                                        "description": "Saturday Delivery",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1500.0
                                            }
                                        ]
                                    },
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 5197.6
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 6780.6
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "SATURDAY_DELIVERY",
                                    "description": "Saturday Delivery",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 1500.0
                                        }
                                    ]
                                },
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 5197.6
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "packageRateDetail": {
                                    "rateType": "PAYOR_LIST_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 6780.6
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 23120.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 4161.6
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 27281.9
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "SATURDAY_DELIVERY",
                                            "description": "Saturday Delivery",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 1500.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 5197.6
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    }
                ],
                "anonymouslyAllowable": true,
                "operationalDetail": {
                    "originLocationIds": [
                        "MAAPD"
                    ],
                    "originLocationNumbers": [
                        0
                    ],
                    "originServiceAreas": [
                        "AA"
                    ],
                    "destinationLocationIds": [
                        "OKCA "
                    ],
                    "destinationLocationNumbers": [
                        0
                    ],
                    "destinationServiceAreas": [
                        "A1"
                    ],
                    "destinationLocationStateOrProvinceCodes": [
                        "OK"
                    ],
                    "deliveryDate": "2025-05-10T13:30:00",
                    "deliveryDay": "SAT",
                    "commitDates": [
                        "2025-05-10T13:30:00"
                    ],
                    "commitDays": [
                        "SAT"
                    ],
                    "ineligibleForMoneyBackGuarantee": false,
                    "astraDescription": "IPE",
                    "originPostalCodes": [
                        "600003"
                    ],
                    "stateOrProvinceCodes": [
                        "TN"
                    ],
                    "countryCodes": [
                        "IN"
                    ],
                    "airportId": "OKC",
                    "serviceCode": "2A",
                    "destinationPostalCodes": [
                        "73102"
                    ]
                },
                "saturdayDelivery": true
            },
            {
                "serviceType": "FEDEX_INTERNATIONAL_PRIORITY",
                "serviceSubOptionDetail": {},
                "serviceName": "FedEx International Priority®",
                "packagingType": "FEDEX_10KG_BOX",
                "commit": {
                    "dateDetail": {
                        "dayOfWeek": "Sat",
                        "time": "17:00",
                        "day": "May 10, 2025",
                        "dayCxsFormat": "May-10-2025"
                    },
                    "commodityName": "BOOKS"
                },
                "commitDetails": [
                    {
                        "dateDetail": {
                            "dayOfWeek": "Sat",
                            "time": "17:00",
                            "day": "May 10, 2025",
                            "dayCxsFormat": "May-10-2025"
                        },
                        "commodityName": "BOOKS"
                    }
                ],
                "customerMessages": [
                    {
                        "code": "customsnote",
                        "message": "<b>Important for Customs:</b> To prevent delays, your shipment must have the following customs documents attached: Commercial Invoice, Certificate of Origin, Shippers Export Declaration. Additional clearance documents may also be required. For more information, please visit our <A HREF=\"https://www.fedex.com/GTM?cntry_code=IN/\">international tools</A> site."
                    },
                    {
                        "code": "postalAwareAdditionalNote",
                        "message": "For countries where zip/postal code or city name can be entered, it is recommended that zip/postal code be entered to obtain a more accurate number of available services."
                    },
                    {
                        "code": "multiweightNote",
                        "message": "Multiweight rated shipment"
                    }
                ],
                "ratedShipmentDetails": [
                    {
                        "rateType": "ACCOUNT",
                        "ratedWeightMethod": "ACTUAL",
                        "effectiveNetDiscount": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 15561.4
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 26097.3
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 22116.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 26097.3
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2P_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 26097.3
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1990.5
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1990.5
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "SATURDAY_DELIVERY",
                                        "description": "Saturday Delivery",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1500.0
                                            }
                                        ]
                                    },
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 4971.9
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 6554.9
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "SATURDAY_DELIVERY",
                                    "description": "Saturday Delivery",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 1500.0
                                        }
                                    ]
                                },
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 4971.9
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "effectiveNetDiscount": {
                                    "currency": "INR",
                                    "amount": 0.0
                                },
                                "packageRateDetail": {
                                    "rateType": "PAYOR_ACCOUNT_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 6554.9
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 22116.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 3981.0
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 26097.3
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "SATURDAY_DELIVERY",
                                            "description": "Saturday Delivery",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 1500.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 4971.9
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    },
                    {
                        "rateType": "LIST",
                        "ratedWeightMethod": "ACTUAL",
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 15561.4
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 26097.3
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 22116.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 26097.3
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2P_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 26097.3
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1990.5
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1990.5
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "SATURDAY_DELIVERY",
                                        "description": "Saturday Delivery",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1500.0
                                            }
                                        ]
                                    },
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 4971.9
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 6554.9
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "SATURDAY_DELIVERY",
                                    "description": "Saturday Delivery",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 1500.0
                                        }
                                    ]
                                },
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 4971.9
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "packageRateDetail": {
                                    "rateType": "PAYOR_LIST_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 6554.9
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 22116.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 3981.0
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 26097.3
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "SATURDAY_DELIVERY",
                                            "description": "Saturday Delivery",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 1500.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 4971.9
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    }
                ],
                "anonymouslyAllowable": true,
                "operationalDetail": {
                    "originLocationIds": [
                        "MAAPD"
                    ],
                    "originLocationNumbers": [
                        0
                    ],
                    "originServiceAreas": [
                        "AA"
                    ],
                    "destinationLocationIds": [
                        "OKCA "
                    ],
                    "destinationLocationNumbers": [
                        0
                    ],
                    "destinationServiceAreas": [
                        "A1"
                    ],
                    "destinationLocationStateOrProvinceCodes": [
                        "OK"
                    ],
                    "deliveryDate": "2025-05-10T17:00:00",
                    "deliveryDay": "SAT",
                    "commitDates": [
                        "2025-05-10T17:00:00"
                    ],
                    "commitDays": [
                        "SAT"
                    ],
                    "ineligibleForMoneyBackGuarantee": false,
                    "astraDescription": "IP EOD",
                    "originPostalCodes": [
                        "600003"
                    ],
                    "stateOrProvinceCodes": [
                        "TN"
                    ],
                    "countryCodes": [
                        "IN"
                    ],
                    "airportId": "OKC",
                    "serviceCode": "2P",
                    "destinationPostalCodes": [
                        "73102"
                    ]
                },
                "saturdayDelivery": true
            },
            {
                "serviceType": "FEDEX_INTERNATIONAL_PRIORITY_EXPRESS",
                "serviceSubOptionDetail": {},
                "serviceName": "FedEx International Priority® Express",
                "packagingType": "FEDEX_10KG_BOX",
                "commit": {
                    "dateDetail": {
                        "dayOfWeek": "Mon",
                        "time": "10:30",
                        "day": "May 12, 2025",
                        "dayCxsFormat": "May-12-2025"
                    },
                    "commodityName": "BOOKS"
                },
                "commitDetails": [
                    {
                        "dateDetail": {
                            "dayOfWeek": "Mon",
                            "time": "10:30",
                            "day": "May 12, 2025",
                            "dayCxsFormat": "May-12-2025"
                        },
                        "commodityName": "BOOKS"
                    }
                ],
                "customerMessages": [
                    {
                        "code": "customsnote",
                        "message": "<b>Important for Customs:</b> To prevent delays, your shipment must have the following customs documents attached: Commercial Invoice, Certificate of Origin, Shippers Export Declaration. Additional clearance documents may also be required. For more information, please visit our <A HREF=\"https://www.fedex.com/GTM?cntry_code=IN/\">international tools</A> site."
                    },
                    {
                        "code": "postalAwareAdditionalNote",
                        "message": "For countries where zip/postal code or city name can be entered, it is recommended that zip/postal code be entered to obtain a more accurate number of available services."
                    },
                    {
                        "code": "multiweightNote",
                        "message": "Multiweight rated shipment"
                    }
                ],
                "ratedShipmentDetails": [
                    {
                        "rateType": "ACCOUNT",
                        "ratedWeightMethod": "ACTUAL",
                        "effectiveNetDiscount": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 16339.7
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 24998.7
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 21185.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 24998.7
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2A_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 24998.7
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1906.7
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1906.7
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 4762.6
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 4845.6
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 4762.6
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "effectiveNetDiscount": {
                                    "currency": "INR",
                                    "amount": 0.0
                                },
                                "packageRateDetail": {
                                    "rateType": "PAYOR_ACCOUNT_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 4845.6
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 21185.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 3813.4
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 24998.7
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 4762.6
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    },
                    {
                        "rateType": "LIST",
                        "ratedWeightMethod": "ACTUAL",
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 16339.7
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 24998.7
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 21185.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 24998.7
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2A_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 24998.7
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1906.7
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1906.7
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 4762.6
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 4845.6
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 4762.6
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "packageRateDetail": {
                                    "rateType": "PAYOR_LIST_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 16339.7
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 4845.6
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 21185.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 3813.4
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 24998.7
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 4762.6
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    }
                ],
                "anonymouslyAllowable": true,
                "operationalDetail": {
                    "originLocationIds": [
                        "MAAPD"
                    ],
                    "originLocationNumbers": [
                        0
                    ],
                    "originServiceAreas": [
                        "AA"
                    ],
                    "destinationLocationIds": [
                        "OKCA "
                    ],
                    "destinationLocationNumbers": [
                        0
                    ],
                    "destinationServiceAreas": [
                        "A1"
                    ],
                    "destinationLocationStateOrProvinceCodes": [
                        "OK"
                    ],
                    "deliveryDate": "2025-05-12T10:30:00",
                    "deliveryDay": "MON",
                    "commitDates": [
                        "2025-05-12T10:30:00"
                    ],
                    "commitDays": [
                        "MON"
                    ],
                    "ineligibleForMoneyBackGuarantee": false,
                    "astraDescription": "IPE",
                    "originPostalCodes": [
                        "600003"
                    ],
                    "stateOrProvinceCodes": [
                        "TN"
                    ],
                    "countryCodes": [
                        "IN"
                    ],
                    "airportId": "OKC",
                    "serviceCode": "2A",
                    "destinationPostalCodes": [
                        "73102"
                    ]
                },
                "saturdayDelivery": false
            },
            {
                "serviceType": "FEDEX_INTERNATIONAL_PRIORITY",
                "serviceSubOptionDetail": {},
                "serviceName": "FedEx International Priority®",
                "packagingType": "FEDEX_10KG_BOX",
                "commit": {
                    "dateDetail": {
                        "dayOfWeek": "Mon",
                        "time": "17:00",
                        "day": "May 12, 2025",
                        "dayCxsFormat": "May-12-2025"
                    },
                    "commodityName": "BOOKS"
                },
                "commitDetails": [
                    {
                        "dateDetail": {
                            "dayOfWeek": "Mon",
                            "time": "17:00",
                            "day": "May 12, 2025",
                            "dayCxsFormat": "May-12-2025"
                        },
                        "commodityName": "BOOKS"
                    }
                ],
                "customerMessages": [
                    {
                        "code": "customsnote",
                        "message": "<b>Important for Customs:</b> To prevent delays, your shipment must have the following customs documents attached: Commercial Invoice, Certificate of Origin, Shippers Export Declaration. Additional clearance documents may also be required. For more information, please visit our <A HREF=\"https://www.fedex.com/GTM?cntry_code=IN/\">international tools</A> site."
                    },
                    {
                        "code": "postalAwareAdditionalNote",
                        "message": "For countries where zip/postal code or city name can be entered, it is recommended that zip/postal code be entered to obtain a more accurate number of available services."
                    },
                    {
                        "code": "multiweightNote",
                        "message": "Multiweight rated shipment"
                    }
                ],
                "ratedShipmentDetails": [
                    {
                        "rateType": "ACCOUNT",
                        "ratedWeightMethod": "ACTUAL",
                        "effectiveNetDiscount": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 15561.4
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 23813.9
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 20181.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 23813.9
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2P_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 23813.9
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1816.3
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1816.3
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 4536.9
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 4619.9
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 4536.9
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "effectiveNetDiscount": {
                                    "currency": "INR",
                                    "amount": 0.0
                                },
                                "packageRateDetail": {
                                    "rateType": "PAYOR_ACCOUNT_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 4619.9
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 20181.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 3632.6
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 23813.9
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 4536.9
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    },
                    {
                        "rateType": "LIST",
                        "ratedWeightMethod": "ACTUAL",
                        "totalBaseCharge": [
                            {
                                "currency": "INR",
                                "amount": 15561.4
                            }
                        ],
                        "totalNetCharge": [
                            {
                                "currency": "INR",
                                "amount": 23813.9
                            }
                        ],
                        "totalVatCharge": [
                            {
                                "amount": 0.0
                            }
                        ],
                        "totalNetFedExCharge": [
                            {
                                "currency": "INR",
                                "amount": 20181.3
                            }
                        ],
                        "totalDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalNetChargeWithDutiesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 23813.9
                            }
                        ],
                        "shipmentLegRateDetails": [
                            {
                                "rateScale": "IN001OFC_2P_FEDEX_10KG_BOX",
                                "totalBaseCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    }
                                ],
                                "totalNetCharge": [
                                    {
                                        "currency": "INR",
                                        "amount": 23813.9
                                    }
                                ],
                                "taxes": [
                                    {
                                        "type": "GST",
                                        "description": "India TN CGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1816.3
                                            }
                                        ]
                                    },
                                    {
                                        "type": "GST",
                                        "description": "India TN SGST",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 1816.3
                                            }
                                        ]
                                    }
                                ],
                                "surcharges": [
                                    {
                                        "type": "FUEL",
                                        "description": "Fuel Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 4536.9
                                            }
                                        ]
                                    },
                                    {
                                        "type": "DEMAND",
                                        "description": "Demand Surcharge",
                                        "amount": [
                                            {
                                                "currency": "INR",
                                                "amount": 83.0
                                            }
                                        ]
                                    }
                                ],
                                "discounts": []
                            }
                        ],
                        "ancillaryFeesAndTaxes": [],
                        "totalDutiesTaxesAndFees": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "totalAncillaryFeesAndTaxes": [
                            {
                                "currency": "INR",
                                "amount": 0.0
                            }
                        ],
                        "shipmentRateDetail": {
                            "rateZone": "C",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 29.0,
                            "totalSurcharges": {
                                "currency": "INR",
                                "amount": 4619.9
                            },
                            "totalFreightDiscount": {
                                "currency": "INR",
                                "amount": 0.0
                            },
                            "freightDiscount": [],
                            "surCharges": [
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 4536.9
                                        }
                                    ]
                                },
                                {
                                    "type": "DEMAND",
                                    "description": "Demand Surcharge",
                                    "amount": [
                                        {
                                            "currency": "INR",
                                            "amount": 83.0
                                        }
                                    ]
                                }
                            ],
                            "totalRateScaleWeight": {
                                "units": "KG",
                                "value": 1.0
                            }
                        },
                        "ratedPackages": [
                            {
                                "groupNumber": 0,
                                "packageRateDetail": {
                                    "rateType": "PAYOR_LIST_SHIPMENT",
                                    "ratedWeightMethod": "ACTUAL",
                                    "baseCharge": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "netFreight": {
                                        "currency": "INR",
                                        "amount": 15561.4
                                    },
                                    "totalSurcharges": {
                                        "currency": "INR",
                                        "amount": 4619.9
                                    },
                                    "netFedExCharge": {
                                        "currency": "INR",
                                        "amount": 20181.3
                                    },
                                    "totalTaxes": {
                                        "currency": "INR",
                                        "amount": 3632.6
                                    },
                                    "netCharge": {
                                        "currency": "INR",
                                        "amount": 23813.9
                                    },
                                    "totalRebates": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "billingWeight": {
                                        "units": "KG",
                                        "value": 0.5
                                    },
                                    "totalFreightDiscounts": {
                                        "currency": "INR",
                                        "amount": 0.0
                                    },
                                    "freightDiscounts": [],
                                    "surcharges": [
                                        {
                                            "type": "FUEL",
                                            "description": "Fuel Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 4536.9
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        },
                                        {
                                            "type": "DEMAND",
                                            "description": "Demand Surcharge",
                                            "amount": [
                                                {
                                                    "currency": "INR",
                                                    "amount": 83.0
                                                }
                                            ],
                                            "level": "SHIPMENT"
                                        }
                                    ],
                                    "rateScaleWeight": {
                                        "units": "KG",
                                        "value": 1.0
                                    }
                                },
                                "sequenceNumber": 1
                            }
                        ]
                    }
                ],
                "anonymouslyAllowable": true,
                "operationalDetail": {
                    "originLocationIds": [
                        "MAAPD"
                    ],
                    "originLocationNumbers": [
                        0
                    ],
                    "originServiceAreas": [
                        "AA"
                    ],
                    "destinationLocationIds": [
                        "OKCA "
                    ],
                    "destinationLocationNumbers": [
                        0
                    ],
                    "destinationServiceAreas": [
                        "A1"
                    ],
                    "destinationLocationStateOrProvinceCodes": [
                        "OK"
                    ],
                    "deliveryDate": "2025-05-12T17:00:00",
                    "deliveryDay": "MON",
                    "commitDates": [
                        "2025-05-12T17:00:00"
                    ],
                    "commitDays": [
                        "MON"
                    ],
                    "ineligibleForMoneyBackGuarantee": false,
                    "astraDescription": "IP EOD",
                    "originPostalCodes": [
                        "600003"
                    ],
                    "stateOrProvinceCodes": [
                        "TN"
                    ],
                    "countryCodes": [
                        "IN"
                    ],
                    "airportId": "OKC",
                    "serviceCode": "2P",
                    "destinationPostalCodes": [
                        "73102"
                    ]
                },
                "saturdayDelivery": false
            }
        ],
        "servicesAvailableAndFiltered": false,
        "quoteDate": "May-07-2025",
        "encoded": false
    }
}
```


# fedex availability api

## request

```bash
curl --location 'https://api.fedex.com/availability/v2/specialserviceoptions' \
--header 'accept: application/json' \
--header 'authorization: Bearer l7bfc00109fa7c4522bca5b1bc6fd6bf0d' \
--header 'content-type: application/json' \
--header 'origin: https://www.fedex.com' \
--header 'referer: https://www.fedex.com/en-in/online/rating.html' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36' \
--header 'x-clientid: MAGR' \
--header 'x-clientversion: 1.0' \
--data '{
    "rateRequestControlParameters": {
        "rateSortOrder": "COMMITASCENDING",
        "returnTransitTimes": true,
        "variableOptions": null,
        "servicesNeededOnRateFailure": false
    },
    "requestedShipment": {
        "shipper": {
            "accountNumber": {
                "key": "a07ef98b62be7bca88e2007220c23084",
                "value": "453673782"
            },
            "address": {
                "city": "Chennai",
                "postalCode": "600003",
                "countryCode": "IN",
                "residential": false,
                "stateOrProvinceCode": "TN"
            }
        },
        "recipients": [
            {
                "address": {
                    "city": "Oklahoma City",
                    "postalCode": "73102",
                    "countryCode": "US",
                    "residential": false,
                    "stateOrProvinceCode": "OK"
                }
            }
        ],
        "shipTimestamp": "2025-05-08",
        "pickupType": "CONTACT_FEDEX_TO_SCHEDULE",
        "packagingType": "FEDEX_10KG_BOX",
        "shippingChargesPayment": {
            "payor": {
                "responsibleParty": {
                    "accountNumber": {
                        "key": "a07ef98b62be7bca88e2007220c23084",
                        "value": "453673782"
                    },
                    "address": {
                        "countryCode": "IN"
                    }
                }
            }
        },
        "blockInsightVisibility": false,
        "edtRequestType": "NONE",
        "rateRequestType": [
            "ACCOUNT",
            "LIST"
        ],
        "requestedPackageLineItems": [
            {
                "groupPackageCount": 1,
                "physicalPackaging": "FEDEX_10KG_BOX",
                "insuredValue": {
                    "currency": "INR",
                    "currencySymbol": null,
                    "amount": 0
                },
                "weight": {
                    "units": "KG",
                    "value": 10
                }
            }
        ],
        "preferredCurrency": "INR",
        "customsClearanceDetail": {
            "commodities": [
                {
                    "name": "NON_DOCUMENTS",
                    "numberOfPieces": 1,
                    "weight": {
                        "units": "KG",
                        "value": 10
                    },
                    "quantity": 1,
                    "quantityUnits": "",
                    "unitPrice": {
                        "currency": "INR",
                        "amount": null
                    },
                    "customsValue": {
                        "currency": "INR",
                        "amount": "1.000",
                        "currencySymbol": ""
                    }
                }
            ]
        }
    },
    "carrierCodes": [
        "FDXG",
        "FDXE"
    ],
    "returnLocalizedDateTime": true,
    "webSiteCountryCode": "IN"
}'
```


## response

{
    "transactionId": "09771cc6-16cb-46d3-bff7-d8b69f6fcd4d",
    "output": {
        "specialServiceOptionsList": [
            {
                "serviceType": "FEDEX_INTERNATIONAL_PRIORITY",
                "carrierCode": "FDXE",
                "shipmentSpecialServicesList": [
                    {
                        "value": "Electronic Trade Documents",
                        "specialServiceType": "ELECTRONIC_TRADE_DOCUMENTS",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "01",
                        "value": "HAL",
                        "specialServiceType": "HOLD_AT_LOCATION",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "40",
                        "value": "Broker Select Option",
                        "specialServiceType": "BROKER_SELECT_OPTION",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "26",
                        "value": "Returns Clearance",
                        "specialServiceType": "RETURNS_CLEARANCE",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "92",
                        "value": "International Traffic In Arms Regulations",
                        "specialServiceType": "INTERNATIONAL_TRAFFIC_IN_ARMS_REGULATIONS",
                        "customerIntegrationMode": "CUSTOM"
                    }
                ],
                "packageSpecialServicesList": [
                    {
                        "code": "93",
                        "value": "Lithium Batteries",
                        "specialServiceType": "BATTERY",
                        "customerIntegrationMode": "CUSTOM"
                    }
                ],
                "returnShipmentList": [
                    "EMAIL_LABEL",
                    "PRINT_RETURN_LABEL"
                ],
                "alertList": [],
                "batteryOptionList": [
                    {
                        "batteryMaterialType": "LITHIUM_METAL",
                        "batteryPackingType": "CONTAINED_IN_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Metal Contained in Equipment (UN3091, PI970)"
                    },
                    {
                        "batteryMaterialType": "LITHIUM_METAL",
                        "batteryPackingType": "PACKED_WITH_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Metal Packed with Equipment (UN3091, PI969)"
                    },
                    {
                        "batteryMaterialType": "LITHIUM_ION",
                        "batteryPackingType": "CONTAINED_IN_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Ion Contained in Equipment (UN3481, PI967)"
                    },
                    {
                        "batteryMaterialType": "LITHIUM_ION",
                        "batteryPackingType": "PACKED_WITH_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Ion Packed with Equipment (UN3481, PI966)"
                    }
                ]
            },
            {
                "serviceType": "FEDEX_INTERNATIONAL_PRIORITY_EXPRESS",
                "carrierCode": "FDXE",
                "shipmentSpecialServicesList": [
                    {
                        "value": "Electronic Trade Documents",
                        "specialServiceType": "ELECTRONIC_TRADE_DOCUMENTS",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "01",
                        "value": "HAL",
                        "specialServiceType": "HOLD_AT_LOCATION",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "40",
                        "value": "Broker Select Option",
                        "specialServiceType": "BROKER_SELECT_OPTION",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "26",
                        "value": "Returns Clearance",
                        "specialServiceType": "RETURNS_CLEARANCE",
                        "customerIntegrationMode": "CUSTOM"
                    },
                    {
                        "code": "92",
                        "value": "International Traffic In Arms Regulations",
                        "specialServiceType": "INTERNATIONAL_TRAFFIC_IN_ARMS_REGULATIONS",
                        "customerIntegrationMode": "CUSTOM"
                    }
                ],
                "packageSpecialServicesList": [
                    {
                        "code": "93",
                        "value": "Lithium Batteries",
                        "specialServiceType": "BATTERY",
                        "customerIntegrationMode": "CUSTOM"
                    }
                ],
                "returnShipmentList": [
                    "EMAIL_LABEL",
                    "PRINT_RETURN_LABEL"
                ],
                "alertList": [],
                "batteryOptionList": [
                    {
                        "batteryMaterialType": "LITHIUM_METAL",
                        "batteryPackingType": "CONTAINED_IN_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Metal Contained in Equipment (UN3091, PI970)"
                    },
                    {
                        "batteryMaterialType": "LITHIUM_METAL",
                        "batteryPackingType": "PACKED_WITH_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Metal Packed with Equipment (UN3091, PI969)"
                    },
                    {
                        "batteryMaterialType": "LITHIUM_ION",
                        "batteryPackingType": "CONTAINED_IN_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Ion Contained in Equipment (UN3481, PI967)"
                    },
                    {
                        "batteryMaterialType": "LITHIUM_ION",
                        "batteryPackingType": "PACKED_WITH_EQUIPMENT",
                        "batteryRegulatoryType": "IATA_SECTION_II",
                        "batterySubtypeDescription": "Ion Packed with Equipment (UN3481, PI966)"
                    }
                ]
            }
        ],
        "signatureOptionsAvailable": true
    }
}