package main

import (
        "fmt"
        "route4me"
)

func main() {
        r4m := route4me.NewClient("11111111111111111111111111111111")


        request := route4me.NewOptimizationNewParams()
        request.Parameters = &route4me.RouteParameters{
                Algorithm_type: route4me.ALGORITHMTYPE_CVRP_TW_SD,
                Distance_unit: route4me.DISTANCEUNIT_MI,
                Device_type: route4me.DEVICETYPE_WEB,
                Optimize: route4me.OPTIMIZE_DISTANCE,
                Metric: route4me.METRIC_GEODESIC,
                Route_max_duration: 86400 * 2,
                Travel_mode: route4me.TRAVELMODE_DRIVING,
                Vehicle_capacity: 50,
                Vehicle_max_distance: 10000,
                Parts: 50,
        }
        request.Addresses = []route4me.Address{
                route4me.Address {
                        Lng: -85.757308,
                        Lat: 38.251698,
                        Is_depot: true,
                        Time: 300,
                        Sequence_no: 0,
                        Address: "455 S 4th St, Louisville, KY 40202",
                },
                route4me.Address {
                        Lng: -85.793846,
                        Lat: 38.141598,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 1,
                        Address: "1604 PARKRIDGE PKWY, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.786514,
                        Lat: 38.202496,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 2,
                        Address: "1407 MCCOY, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.774864,
                        Lat: 38.178844,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 3,
                        Address: "4805 BELLEVUE AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.821121,
                        Lat: 38.248684,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 4,
                        Address: "730 CECIL AVENUE, Louisville, KY, 40211",
                },
                route4me.Address {
                        Lng: -85.800034,
                        Lat: 38.251923,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 5,
                        Address: "650 SOUTH 29TH ST UNIT 315, Louisville, KY, 40211",
                },
                route4me.Address {
                        Lng: -85.824638,
                        Lat: 38.176067,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 6,
                        Address: "4629 HILLSIDE DRIVE, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.775558,
                        Lat: 38.179806,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 7,
                        Address: "4738 BELLEVUE AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.815094,
                        Lat: 38.259335,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 8,
                        Address: "318 SO. 39TH STREET, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.785118,
                        Lat: 38.179253,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 9,
                        Address: "1324 BLUEGRASS AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.792854,
                        Lat: 38.162472,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 10,
                        Address: "7305 ROYAL WOODS DR, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.783966,
                        Lat: 38.229584,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 11,
                        Address: "1661 W HILL ST, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.822594,
                        Lat: 38.210606,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 12,
                        Address: "3222 KINGSWOOD WAY, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.796783,
                        Lat: 38.153767,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 13,
                        Address: "1922 PALATKA RD, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.796852,
                        Lat: 38.235847,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 14,
                        Address: "1314 SOUTH 26TH STREET, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.789032,
                        Lat: 38.218662,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 15,
                        Address: "2135 MCCLOSKEY AVENUE, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.781387,
                        Lat: 38.206154,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 16,
                        Address: "1409 PHYLLIS AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.839149,
                        Lat: 38.187511,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 17,
                        Address: "4504 SUNFLOWER AVE, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.795059,
                        Lat: 38.241405,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 18,
                        Address: "2512 GREENWOOD AVE, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.863319,
                        Lat: 38.166065,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 19,
                        Address: "5500 WILKE FARM AVE, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.786201,
                        Lat: 38.193283,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 20,
                        Address: "3640 LENTZ AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.780037,
                        Lat: 38.17952,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 21,
                        Address: "1020 BLUEGRASS AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.814156,
                        Lat: 38.26498,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 22,
                        Address: "123 NORTH 40TH ST, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.802867,
                        Lat: 38.151072,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 23,
                        Address: "7315 ST ANDREWS WOODS CIRCLE UNIT 104, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.849937,
                        Lat: 38.182594,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 24,
                        Address: "3210 POPLAR VIEW DR, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.811447,
                        Lat: 38.1754,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 25,
                        Address: "4519 LOUANE WAY, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.798279,
                        Lat: 38.161839,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 26,
                        Address: "6812 MANSLICK RD, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.788353,
                        Lat: 38.172031,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 27,
                        Address: "1524 HUNTOON AVENUE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.779816,
                        Lat: 38.209663,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 28,
                        Address: "1307 LARCHMONT AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.791962,
                        Lat: 38.26844,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 29,
                        Address: "434 N 26TH STREET #2, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.80629,
                        Lat: 38.250397,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 30,
                        Address: "678 WESTLAWN ST, Louisville, KY, 40211",
                },
                route4me.Address {
                        Lng: -85.790421,
                        Lat: 38.248882,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 31,
                        Address: "2308 W BROADWAY, Louisville, KY, 40211",
                },
                route4me.Address {
                        Lng: -85.794257,
                        Lat: 38.233579,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 32,
                        Address: "2332 WOODLAND AVE, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.783928,
                        Lat: 38.239697,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 33,
                        Address: "1706 WEST ST. CATHERINE, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.792397,
                        Lat: 38.216465,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 34,
                        Address: "1699 WATHEN LN, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.831787,
                        Lat: 38.186245,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 35,
                        Address: "2416 SUNSHINE WAY, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.798355,
                        Lat: 38.158466,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 36,
                        Address: "6925 MANSLICK RD, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.785082,
                        Lat: 38.212438,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 37,
                        Address: "2707 7TH ST, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.826668,
                        Lat: 38.179394,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 38,
                        Address: "2014 KENDALL LN, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.812012,
                        Lat: 38.273354,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 39,
                        Address: "612 N 39TH ST, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.786781,
                        Lat: 38.261703,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 40,
                        Address: "2215 ROWAN ST, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.78653,
                        Lat: 38.241611,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 41,
                        Address: "1826 W. KENTUCKY ST, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.796211,
                        Lat: 38.224716,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 42,
                        Address: "1810 GREGG AVE, Louisville, KY, 40210",
                },
                route4me.Address {
                        Lng: -85.825836,
                        Lat: 38.191753,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 43,
                        Address: "4103 BURRRELL DRIVE, Louisville, KY, 40216",
                },
                route4me.Address {
                        Lng: -85.823463,
                        Lat: 38.259903,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 44,
                        Address: "359 SOUTHWESTERN PKWY, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.792109,
                        Lat: 38.252781,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 45,
                        Address: "2407 W CHESTNUT ST, Louisville, KY, 40211",
                },
                route4me.Address {
                        Lng: -85.786658,
                        Lat: 38.257616,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 46,
                        Address: "225 S 22ND ST, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.786072,
                        Lat: 38.202122,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 47,
                        Address: "1404 MCCOY AVE, Louisville, KY, 40215",
                },
                route4me.Address {
                        Lng: -85.799438,
                        Lat: 38.270061,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 48,
                        Address: "117 FOUNT LANDING CT, Louisville, KY, 40212",
                },
                route4me.Address {
                        Lng: -85.7798,
                        Lat: 38.145851,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 49,
                        Address: "5504 SHOREWOOD DRIVE, Louisville, KY, 40214",
                },
                route4me.Address {
                        Lng: -85.780251,
                        Lat: 38.211025,
                        Is_depot: false,
                        Time: 300,
                        Sequence_no: 50,
                        Address: "1406 CENTRAL AVE, Louisville, KY, 40208",
                },
        }

        response, exception, err := r4m.NewOptimization(request)
        if err != nil {
                fmt.Print(err)
                return
        }
        if exception != nil {
                for _, error := range exception.Errors {
                        fmt.Print(error)
                }
                return
        }
        fmt.Printf("%+v", response)
}
