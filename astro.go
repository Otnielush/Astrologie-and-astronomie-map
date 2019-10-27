package astro

// http://cosinekitty.com/solar_system.html

import (
	"fmt"
	"math"
)

const (
	CentX, CentY = 400, 300
	T            = 6.28318530718
	Pi           = T / 2
)

//  КМ * 10^7
const (
	MerDis    = float64(5.7910000)
	VenDis    = float64(10.8200000)
	EarDis    = float64(14.9600000)
	MarDis    = float64(22.7900000)
	JupDis    = float64(77.8500000)
	SatDis    = float64(143.4000000)
	UraDis    = float64(287.1000000)
	NepDis    = float64(449.5000000)
	MoonDis   = float64(0.0384400)
	ZodiakDis = float64(450) // Просто линия
	AU        = float64(14.95978707)
)
const (
	MerDur  = 87.9365
	VenDur  = 224.701
	EarDur  = 365.25
	MarDur  = 687.35
	JupDur  = 4344.0
	SatDur  = 10670.0
	UraDur  = 31000.0
	NepDur  = 60186.0
	MoonDur = 27.324151
	YearDur = 365.2515
)

type planetPos struct {
	X, Y int32
}

// POSITIONS ON 2000 IN °

const (
	startMer  = 249.416 / 360.0
	startVen  = 180.03 / 360.0
	startEar  = (98.76) / 360.0
	startMar  = 358.45 / 360.0
	startJup  = 36.15 / 360.0
	startSat  = 45.65 / 360.0
	startUra  = 316.383 / 360.0
	startNep  = 303.9 / 360.0
	startMoon = 98.916 / 360.0
)

var gradVzodiak = map[float64]int{
	0.000:  1,
	8.333:  2,
	16.667: 3,
	25.000: 4,
	33.333: 5,
	41.667: 6,
	50.000: 7,
	58.333: 8,
	66.667: 9,
	75.000: 10,
	83.333: 11,
	91.667: 12,
}
var nameZod = map[int]string{
	1:  "Овен",
	2:  "Телец",
	3:  "Близнецы",
	4:  "Рак",
	5:  "Лев",
	6:  "Дева",
	7:  "Весы",
	8:  "Скорпион",
	9:  "Стрелец",
	10: "Козерог",
	11: "Водолей",
	12: "Рыбы",
}

var hMonth = map[string]int{"Nisan": 1, "Iyyar": 2, "Sivan": 3, "Tamuz": 4, "Av": 5, "Elul": 6, "Tishrei": 7, "Cheshvan": 8, "Kislev": 9, "Tevet": 10, "Shvat": 11, "Adar1": 12, "Adar2": 13}

type planetInfo struct {
	Name         string
	Zodiak       int
	Angle_Helio  float64
	year         float64
	radius       float64
	startPos     float64
	Angle_Geo    float64
	Angle_zodiak float64
	Xecl         float64
	Yecl         float64
	Zecl         float64
	pos          possPlanet
}
type possPlanet struct {
	N         float64 // longitude of the ascending node (градус от перилиона) элипс
	N_per_day float64
	i         float64 // inclination to the ecliptic
	i_per_day float64
	w         float64 // longitude of perihelion
	w_per_day float64
	a         float64 // mean distance, a.u.
	a_per_day float64
	e         float64 // eccentricity 0=circle, 0-1=ellipse, 1=parabola
	e_per_day float64
	M         float64 // mean anomaly  0 at perihelion
	M_per_day float64
	L         float64 // w + m   mean longitude
	v         float64 // true anomaly (angle between position and perihelion)
	E         float64 // eccentric anomaly
}
type sunSystem struct {
	Obj [10]planetInfo
}

var planets = [10]planetInfo{
	{Name: "Sun", Xecl: CentX, Yecl: CentY, year: 0, pos: possPlanet{N: 0, i: 0, w: 282.9404, w_per_day: 4.70935E-5, a: 1.0, e: 0.016709, e_per_day: -1.151E-9, M: 356.0470, M_per_day: 0.9856002585}},
	{Name: "Mercury", startPos: startMer, year: MerDur, radius: MerDis, pos: possPlanet{N: 48.3313, N_per_day: 3.24587E-5, i: 7.0047, i_per_day: 5.00E-8, w: 29.1241, w_per_day: 1.01444E-5, a: 0.387098, e: 0.205635, e_per_day: 5.59E-10, M: 168.6562, M_per_day: 4.0923344368}},
	{Name: "Venus", startPos: startVen, year: VenDur, radius: VenDis, pos: possPlanet{N: 76.6799, N_per_day: 2.46590E-5, i: 3.3946, i_per_day: 2.75E-8, w: 54.8910, w_per_day: 1.38374E-5, a: 0.723330, e: 0.006773, e_per_day: -1.302E-9, M: 48.0052, M_per_day: 1.6021302244}},
	{Name: "Earth", startPos: startEar, year: EarDur, radius: EarDis},
	{Name: "Mars", startPos: startMar, year: MarDur, radius: MarDis, pos: possPlanet{N: 49.5574, N_per_day: 2.11081E-5, i: 1.8497, i_per_day: -1.78E-8, w: 286.5016, w_per_day: 2.92961E-5, a: 1.523688, e: 0.093405, e_per_day: 2.516E-9, M: 18.6021, M_per_day: 0.5240207766}},
	{Name: "Jupiter", startPos: startJup, year: JupDur, radius: JupDis, pos: possPlanet{N: 100.4542, N_per_day: 2.76854E-5, i: 1.3030, i_per_day: -1.557E-7, w: 273.8777, w_per_day: 1.64505E-5, a: 5.20256, e: 0.048498, e_per_day: 4.469E-9, M: 19.8950, M_per_day: 0.0830853001}},
	{Name: "Saturn", startPos: startSat, year: SatDur, radius: SatDis, pos: possPlanet{N: 113.6634, N_per_day: 2.38980E-5, i: 2.4886, i_per_day: -1.081E-7, w: 339.3939, w_per_day: 2.97661E-5, a: 9.55475, e: 0.055546, e_per_day: -9.499E-9, M: 316.9670, M_per_day: 0.0334442282}},
	{Name: "Uranus", startPos: startUra, year: UraDur, radius: UraDis, pos: possPlanet{N: 74.0005, N_per_day: 1.3978E-5, i: 0.7733, i_per_day: 1.9E-8, w: 96.6612, w_per_day: 3.0565E-5, a: 19.18171, a_per_day: -1.55E-8, e: 0.047318, e_per_day: 7.45E-9, M: 142.5905, M_per_day: 0.011725806}},
	{Name: "Neptune", startPos: startNep, year: NepDur, radius: NepDis, pos: possPlanet{N: 131.7806, N_per_day: 3.0173E-5, i: 1.7700, i_per_day: -2.55E-7, w: 272.8461, w_per_day: -6.027E-6, a: 30.05826, a_per_day: 3.313E-8, e: 0.008606, e_per_day: 2.15E-9, M: 260.2471, M_per_day: 0.005995147}},
	{Name: "Moon", startPos: startMoon, year: MoonDur, radius: MoonDis, pos: possPlanet{N: 125.1228, N_per_day: -0.0529538083, i: 5.1454, w: 318.0634, w_per_day: 0.1643573223, a: 60.2666, e: 0.0549, M: 115.3654, M_per_day: 13.0649929509}},
}

// var SS = sunSystem{Obj: planets}

var p = fmt.Println
var pf = fmt.Printf

// NEW

const (
	oblecl2000  = 23.4393 // obliquity of the ecliptic
	obl_per_day = 3.563E-7
)

func CalcDays(d, m, y int) float64 {
	return float64(367*y - 7*(y+(m+9)/12)/4 + 275*m/9 + d - 730530)
}

func (pl *planetInfo) calcPos(day float64) {

	N := pl.pos.N + pl.pos.N_per_day*day
	i := pl.pos.i + pl.pos.i_per_day*day
	// p("N", N, " i", i)
	w := pl.pos.w + pl.pos.w_per_day*day
	rev(&w)
	a := pl.pos.a // для луны это в диаметрах Земли
	e := pl.pos.e + pl.pos.e_per_day*day
	M := pl.pos.M + pl.pos.M_per_day*day
	rev(&M)
	// p("w", w, "  a", a, "  e", e)

	// oblecl := oblecl2000 + obl_per_day*day
	pl.pos.L = w + M
	rev(&pl.pos.L)
	// p("M", M, "  L", pl.pos.L)

	pl.pos.E = M + (180/Pi)*e*math.Sin((M*Pi/180))*(1+e*math.Cos((M*Pi/180)))
	// if e > 0.055 {
	pl.pos.E = pl.pos.E - (pl.pos.E-e*(180/Pi)*math.Sin(pl.pos.E*Pi/180)-M)/(1-e*math.Cos(pl.pos.E*Pi/180))
	// }
	// p("E", pl.pos.E)

	x := a * (math.Cos(pl.pos.E*Pi/180) - e)
	y := a * (math.Sin(pl.pos.E*Pi/180) * math.Sqrt(1-e*e))
	// p("x", x, "  y", y)

	r := math.Sqrt(x*x + y*y)
	v := math.Atan2(y, x) / T * 360
	rev(&v)

	// p("r", r, "  v", v)

	lon := v + w
	rev(&lon)
	// p("lon", lon)

	// possition in space and Zodiak
	if pl.Name != "Moon" {

		xh := a * (math.Cos(N*Pi/180)*math.Cos((v*Pi/180+w*Pi/180)) - math.Sin(N*Pi/180)*math.Sin((v*Pi/180+w*Pi/180))*math.Cos(i*Pi/180))
		yh := a * (math.Sin(N*Pi/180)*math.Cos((v*Pi/180+w*Pi/180)) + math.Cos(N*Pi/180)*math.Sin((v*Pi/180+w*Pi/180))*math.Cos(i*Pi/180))
		zh := a * (math.Sin((v*Pi/180 + w*Pi/180)) * math.Sin(i*Pi/180))
		// p("xh", xh, " yh", yh, " zh", zh)

		lonEcl := math.Atan2(yh, xh) / T * 360 // Это то что нам нужно
		rev(&lonEcl)
		// latEcl := math.Atan2(zh, math.Sqrt(xh*xh+yh*yh)) / T * 360

		pl.Xecl = xh
		pl.Yecl = yh
		pl.Zecl = zh
		pl.Angle_Helio = lonEcl

		if pl.Name == "Sun" {
			pl.Angle_Geo = lonEcl
			pl.Angle_zodiak = lonEcl
			for pl.Angle_zodiak > 30 {
				pl.Angle_zodiak -= 30
			}
			pl.Zodiak = int(lonEcl/30) + 1
		}

		// p("lonEcl", lonEcl, " latEcl", latEcl, " zod", nameZod[pl.Zodiak])

	} else {
		xeclip := r * (math.Cos(N*Pi/180)*math.Cos((v*Pi/180+w*Pi/180)) - math.Sin(N*Pi/180)*math.Sin((v*Pi/180+w*Pi/180))*math.Cos(i*Pi/180))
		yeclip := r * (math.Sin(N*Pi/180)*math.Cos((v*Pi/180+w*Pi/180)) + math.Cos(N*Pi/180)*math.Sin((v*Pi/180+w*Pi/180))*math.Cos(i*Pi/180))
		zeclip := r * (math.Sin((v*Pi/180 + w*Pi/180)) * math.Sin(i*Pi/180))
		// p("xeclip", xeclip, " yeclip", yeclip, " zeclip", zeclip)

		long := math.Atan2(yeclip, xeclip) / T * 360 // Это то что нам нужно
		// long -= 3
		rev(&long)
		// lat := math.Atan2(zeclip, math.Sqrt(xeclip*xeclip+yeclip*yeclip)) / T * 360
		r = math.Sqrt(xeclip*xeclip + yeclip*yeclip + zeclip*zeclip)

		pl.Xecl = xeclip
		pl.Yecl = yeclip
		pl.Zecl = zeclip
		pl.Angle_Helio = long

		pl.Angle_Geo = long
		pl.Angle_zodiak = long
		for pl.Angle_zodiak > 30 {
			pl.Angle_zodiak -= 30
		}
		pl.Zodiak = int(long/30) + 1

		// p("long", long, " lat", lat, " Zod", nameZod[pl.Zodiak])
	}
}
func (pl *planetInfo) calcGeoPos(xh, yh, zh float64) {
	xg := pl.Xecl + xh
	yg := pl.Yecl + yh
	// zg := pl.Zecl + zh

	pl.Angle_Geo = math.Atan2(yg, xg) / T * 360
	rev(&pl.Angle_Geo)
	pl.Angle_zodiak = pl.Angle_Geo
	for pl.Angle_zodiak > 30 {
		pl.Angle_zodiak -= 30
	}
	pl.Zodiak = int(pl.Angle_Geo/30) + 1
	// p("zod", pl.Zodiak)
}

func (ss *sunSystem) Calculate(day float64) {
	for i := 0; i < 10; i++ {
		ss.Obj[i].calcPos(day)
	}
	for i := 1; i < 9; i++ { // С луной пока не решил
		ss.Obj[i].calcGeoPos(ss.Obj[0].Xecl, ss.Obj[0].Yecl, ss.Obj[0].Zecl)
	}
}
func NewAstro() *sunSystem {
	return &sunSystem{Obj: planets}
}

// func main() {

// 	var y, m, D int
// 	D = 9
// 	m = 9
// 	y = 1990

// 	d := CalcDays(D, m, y)

// 	SS.Calculate(d)

// 	for i := 0; i < 10; i++ {
// 		pf("%s (H) %.2f°,(G) %s %.2f°\n", SS.Obj[i].Name, SS.Obj[i].Angle_Helio, nameZod[SS.Obj[i].Zodiak], SS.Obj[i].Angle_zodiak)
// 	}
// }

func rev(x *float64) {
	var rv float64
	rv = *x - math.Floor(*x/360.0)*360.0
	if rv < 0.0 {
		rv += 360.0
	}
	*x = rv
}
