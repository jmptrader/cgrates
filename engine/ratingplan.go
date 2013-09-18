/*
Rating system designed to be used in VoIP Carriers World
Copyright (C) 2013 ITsysCOM

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package engine

import (
	"github.com/cgrates/cgrates/cache2go"
	"time"
)

/*
The struture that is saved to storage.
*/
type RatingPlan struct {
	ActivationTime time.Time
	RateIntervals  RateIntervalList
}

type xCachedRatingPlans struct {
	destPrefix string
	aps        []*RatingPlan
	*cache2go.XEntry
}

/*
Adds one ore more intervals to the internal interval list only if it is not allready in the list.
*/
func (ap *RatingPlan) AddRateInterval(is ...*RateInterval) {
	for _, i := range is {
		found := false
		for _, ei := range ap.RateIntervals {
			if i.Equal(ei) {
				(&ei.Rates).AddRate(i.Rates...)
				found = true
				break
			}
		}
		if !found {
			ap.RateIntervals = append(ap.RateIntervals, i)
		}
	}
}

func (ap *RatingPlan) Equal(o *RatingPlan) bool {
	return ap.ActivationTime == o.ActivationTime
}