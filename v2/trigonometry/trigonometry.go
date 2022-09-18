package trigonometry

import (
	"math"
)

// Angle represents a DMS angle type.
type Angle struct {
	Degree float64
	Minute float64
	Second float64
}

const (
	ZeroAngle     = 0.0 // degrees
	RightAngle    = 90.0
	StraightAngle = 180.0
	CompleteAngle = 360.0
)

// DegreeToRadian returns angle in radian from given degree d.
func DegreeToRadian(d float64) float64 {
	return (d * math.Pi) / 180.0
}

// RadianToDegree returns angle in degree from given radian r.
func RadianToDegree(r float64) float64 {
	return (r * 180.0) / math.Pi
}

// DegreeToDMS returns a DMS angle from given degree d.
func DegreeToDMS(d float64) Angle {
	degreePart := int(d)

	remainingDegreePart := d - float64(degreePart)
	if remainingDegreePart == 0.0 {
		return Angle{float64(degreePart), 0, 0}
	}

	minutePartTemp := remainingDegreePart * 60
	minutePart := int(minutePartTemp)
	remainingMinutePart := minutePartTemp - float64(minutePart)
	if remainingMinutePart == 0.0 {
		return Angle{float64(degreePart), float64(minutePart), 0}
	}

	secondPart := int(remainingMinutePart * 60)

	return Angle{float64(degreePart), float64(minutePart), float64(secondPart)}
}

// RadianToDMS returns a DMS angle from given radian r.
func RadianToDMS(r float64) Angle {
	d := RadianToDegree(r)
	return DegreeToDMS(d)
}

// DMSToDegree returns angle in degree from the DMS angle.
func (a *Angle) DMSToDegree() float64 {
	d := a.Degree
	dm := a.Minute / 60
	ds := a.Second / 60 / 60
	return d + dm + ds
}

// DMSToRadian returns angle in radian from the DMS angle.
func (a *Angle) DMSToRadian() float64 {
	d := a.DMSToDegree()
	return DegreeToRadian(d)
}

// IsAcuteAngle returns true if DMS angle is acute angle.
func (a *Angle) IsAcuteAngle() bool {
	angle := a.DMSToDegree()

	if angle > 0 && angle < RightAngle {
		return true
	}

	return false
}

// IsObtuseAngle returns true if DMS angle is obtuse angle.
func (a *Angle) IsObtuseAngle() bool {
	angle := a.DMSToDegree()

	if angle > RightAngle && angle < StraightAngle {
		return true
	}

	return false
}

// IsRightAngle returns true if DMS angle is right angle.
func (a *Angle) IsRightAngle() bool {
	angle := a.DMSToDegree()

	if angle == RightAngle {
		return true
	}

	return false
}

// IsCompleteAngle returns true if DMS angle is complete angle.
func (a *Angle) IsCompleteAngle() bool {
	angle := a.DMSToDegree()

	if angle == 0.0 {
		return false
	}

	if angle == CompleteAngle {
		return true
	}

	if math.Mod(angle, CompleteAngle) == 0.0 {
		return true
	}

	return false
}
