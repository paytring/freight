# auth

## request
```bash
curl --location 'https://apis-sandbox.fedex.com/oauth/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--header 'Cookie: ak_bmsc=8B7E48F9072DA3B02A6072E2E77A850A~000000000000000000000000000000~YAAQtkLHFzEEnJGWAQAAtFeLnxuzfdmZlbGMyJIHjbQO66aR0rO9r3j7cBjR5gmgFuBzqB3SB0uztVNBUaZMYraFoG6Cf7T4QAvPe8wrisWWUtl1tRDFBO6W151ffCDcMMei4v5YPM7HMWGDfHc+VQSdZ2LMYjKp6WPw+aAFMpDpktd8M6mVikR7y/fHUQtEWeB+0PeEnGQjZEfmIk8ihwMfvyFkMN4tWYyXqt6Hrxkbb5iJqRIiV218MUAWkvI5QLuNEyi3aLk1kaorlOjRX6rXgZ8H4/w1RJLKD4H03rdAaHe7ecWB6uDRehwqAYGmEPbR1rS+Q5Gtm40kNOQLCrbveRsV7Xzar31y; bm_sv=2DE454D50E4C143FB4664DC5DF231BB5~YAAQldcLF3Tm4XKWAQAAUEuVnxsDcQYJJNhhIJJLeOmy+Og6IuhawgUeigayg9wcP6qTcvJ7k4RuCW5/011n3GD18AzlohfuUD3FBi7NTzOzU4qH0Rw5H0ZReEZIJuF2e3+aD1Bt9oKKNxiWN5vSGGXntqfUL7kC/xotXQK8Ra2KvHdAQcn/uLEWGVPHKKj0v8t3T/f2gXUzpzwBCoAZ3WNKAFgQvVBkrfvQjiU/JyMrXxY89kOAUso7Eik6ULU=~1; fdx_bman=096286ced3e0502d5ce9ea56eb5cf3ca' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'client_id=l7bc7d13b094664ad9a08e01ce9a20fe4c' \
--data-urlencode 'client_secret=671b15a5f28f48fcb87228d92934f060'

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
curl --location 'https://apis-sandbox.fedex.com/rate/v1/rates/quotes' \
--header 'X-locale: en_US' \
--header 'Content-Type: application/json' \
--header 'authorization: Bearer $token' \
--data '{
    "accountNumber": {
        "value": "123456789"
    },
    "requestedShipment": {
        "shipper": {
            "address": {
                "postalCode": "m1m1m1",
                "countryCode": "CA"
            }
        },
        "recipient": {
            "address": {
                "postalCode": "38116",
                "countryCode": "US",
                "residential": true
            }
        },
        
        "pickupType": "CONTACT_FEDEX_TO_SCHEDULE",
        "packagingType": "FEDEX_25KG_BOX",
        "rateRequestType": [
            "ACCOUNT",
            "LIST"
        ],
        "requestedPackageLineItems": [
            {
                "subPackagingType": "CASE",
                "declaredValue": {
                    "amount": 100,
                    "currency": "CAD"
                },
                "weight": {
                    "units": "LB",
                    "value": 22
                },
                "dimensions": {
                    "length": 10,
                    "width": 8,
                    "height": 2,
                    "units": "IN"
                }
            }
        ]
    }
}'
```

## response
```json
{
    "transactionId": "APIF_SV_RATC_TxID58a125b8-47bf-4838-8530-43c8c55d4ab6",
    "output": {
        "alerts": [
            {
                "code": "VIRTUAL.RESPONSE",
                "message": "This is a Virtual Response.",
                "alertType": "NOTE"
            },
            {
                "code": "ORIGIN.STATEORPROVINCECODE.CHANGED",
                "message": "The origin state/province code has been changed.",
                "alertType": "NOTE"
            },
            {
                "code": "DESTINATION.STATEORPROVINCECODE.CHANGED",
                "message": "The destination state/province code has been changed.",
                "alertType": "NOTE"
            }
        ],
        "rateReplyDetails": [
            {
                "serviceType": "FEDEX_INTERNATIONAL_PRIORITY",
                "serviceName": "FedEx International PriorityÂ®",
                "packagingType": "FEDEX_25KG_BOX",
                "customerMessages": [
                    {
                        "code": "SERVICE.TYPE.INTERNATIONAL.MESSAGE",
                        "message": "Rate does not include duties & taxes, clearance entry fees or other import fees.  The payor of duties/taxes/fees will be responsible for any applicable Clearance Entry Fees."
                    }
                ],
                "ratedShipmentDetails": [
                    {
                        "rateType": "ACCOUNT",
                        "ratedWeightMethod": "ACTUAL",
                        "totalDiscounts": 0.0,
                        "totalBaseCharge": 339.99,
                        "totalNetCharge": 416.75,
                        "totalVatCharge": 0.0,
                        "totalNetFedExCharge": 416.75,
                        "totalDutiesAndTaxes": 0.0,
                        "totalNetChargeWithDutiesAndTaxes": 416.75,
                        "totalDutiesTaxesAndFees": 0.0,
                        "totalAncillaryFeesAndTaxes": 0.0,
                        "shipmentRateDetail": {
                            "rateZone": "A",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 20.0,
                            "totalSurcharges": 76.76,
                            "totalFreightDiscount": 0.0,
                            "surCharges": [
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": 69.46
                                },
                                {
                                    "type": "RESIDENTIAL_DELIVERY",
                                    "description": "Residential delivery surcharge",
                                    "amount": 5.8
                                },
                                {
                                    "type": "PEAK",
                                    "description": "Peak Surcharge",
                                    "amount": 1.5
                                }
                            ],
                            "pricingCode": "",
                            "currencyExchangeRate": {
                                "fromCurrency": "USD",
                                "intoCurrency": "USD",
                                "rate": 1.0
                            },
                            "totalBillingWeight": {
                                "units": "LB",
                                "value": 15.0
                            },
                            "dimDivisorType": "COUNTRY",
                            "currency": "USD",
                            "rateScale": "US001DFA_2P_FEDEX_25KG_BOX",
                            "totalRateScaleWeight": {
                                "units": "LB",
                                "value": 15.0
                            }
                        },
                        "currency": "USD"
                    },
                    {
                        "rateType": "LIST",
                        "ratedWeightMethod": "ACTUAL",
                        "totalDiscounts": 0.0,
                        "totalBaseCharge": 339.99,
                        "totalNetCharge": 416.75,
                        "totalVatCharge": 0.0,
                        "totalNetFedExCharge": 416.75,
                        "totalDutiesAndTaxes": 0.0,
                        "totalNetChargeWithDutiesAndTaxes": 416.75,
                        "totalDutiesTaxesAndFees": 0.0,
                        "totalAncillaryFeesAndTaxes": 0.0,
                        "shipmentRateDetail": {
                            "rateZone": "A",
                            "dimDivisor": 139,
                            "fuelSurchargePercent": 20.0,
                            "totalSurcharges": 76.76,
                            "totalFreightDiscount": 0.0,
                            "surCharges": [
                                {
                                    "type": "FUEL",
                                    "description": "Fuel Surcharge",
                                    "amount": 69.46
                                },
                                {
                                    "type": "RESIDENTIAL_DELIVERY",
                                    "description": "Residential delivery surcharge",
                                    "amount": 5.8
                                },
                                {
                                    "type": "PEAK",
                                    "description": "Peak Surcharge",
                                    "amount": 1.5
                                }
                            ],
                            "pricingCode": "",
                            "currencyExchangeRate": {
                                "fromCurrency": "USD",
                                "intoCurrency": "USD",
                                "rate": 1.0
                            },
                            "totalBillingWeight": {
                                "units": "LB",
                                "value": 15.0
                            },
                            "dimDivisorType": "COUNTRY",
                            "currency": "USD",
                            "rateScale": "US001DFA_2P_FEDEX_25KG_BOX",
                            "totalRateScaleWeight": {
                                "units": "LB",
                                "value": 15.0
                            }
                        },
                        "currency": "USD"
                    }
                ],
                "operationalDetail": {
                    "ineligibleForMoneyBackGuarantee": false,
                    "astraDescription": "IP EOD",
                    "airportId": "MEM",
                    "serviceCode": "2P"
                },
                "signatureOptionType": "NO_SIGNATURE_REQUIRED",
                "serviceDescription": {
                    "serviceId": "EP1000000300",
                    "serviceType": "FEDEX_INTERNATIONAL_PRIORITY",
                    "code": "2P",
                    "names": [
                        {
                            "type": "long",
                            "encoding": "utf-8",
                            "value": "FedEx International PriorityÂ®"
                        },
                        {
                            "type": "long",
                            "encoding": "ascii",
                            "value": "FedEx International Priority"
                        },
                        {
                            "type": "medium",
                            "encoding": "utf-8",
                            "value": "FedEx International Priority"
                        },
                        {
                            "type": "medium",
                            "encoding": "ascii",
                            "value": "FedEx International Priority"
                        },
                        {
                            "type": "short",
                            "encoding": "utf-8",
                            "value": "IPED"
                        },
                        {
                            "type": "short",
                            "encoding": "ascii",
                            "value": "IPED"
                        },
                        {
                            "type": "abbrv",
                            "encoding": "ascii",
                            "value": "OA"
                        }
                    ],
                    "serviceCategory": "parcel",
                    "description": "International Priority EOD (IP EOD)",
                    "astraDescription": "IP EOD"
                }
            }
        ],
        "quoteDate": "2025-05-04",
        "encoded": false
    }
}
````