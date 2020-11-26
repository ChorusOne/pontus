package snapshot

import (
	"log"
	"math/big"

	"github.com/oleiade/reflections"
)

var integrityTests = [](func(Snapshot, SystemSnapshot) bool){
	skaleTokenSupplyIntegrity,
}

const (
	pass = "PASS"
	fail = "FAIL"
)

func CheckIntegrity(as Snapshot, ss SystemSnapshot) bool {

	for _, testFunction := range integrityTests {
		testFunction(as, ss)
	}
	return true
}

func skaleTokenSupplyIntegrity(as Snapshot, ss SystemSnapshot) bool {
	sumOf := snapshotColumnSum(as, "SkaleTokenBalance")
	systemMetric := ss.SkaleTokenSupply
	message := "Integrity Test - SkaleTokenSupply - %s - Summation = %s, SystemWideValue = %s \n"

	if sumOf.Cmp(systemMetric) != 0 {
		log.Printf(message, fail, sumOf, systemMetric)
		log.Printf("Delta SystemWideValue - Summation = %s", big.NewInt(0).Sub(systemMetric, sumOf))
		return false
	}

	log.Printf(message, pass, sumOf, systemMetric)
	return true

}

func snapshotColumnSum(as Snapshot, columnName string) *big.Int {
	sum := big.NewInt(0)
	for address := range as {
		val, _ := reflections.GetField(as[address], columnName)
		//Asserting val as *big.Int
		sum.Add(sum, val.(*big.Int))
	}
	return sum
}
