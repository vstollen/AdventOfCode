package fabric

import java.io.File

fun main(args : Array<String>) {

    class ClothClaim(claim: String) {
        val id = claim.substringAfter('#').substringBefore(" @").toInt()

        val sideDistance = claim.substringAfter("@ ").substringBefore(',').toInt()
        val topDistance = claim.substringAfter(',').substringBefore(':').toInt()

        val width = claim.substringAfter(": ").substringBefore('x').toInt()
        val height = claim.substringAfter('x').toInt()

        var overlapping = false
    }

    val input = File("input/day3_fabric.txt")
    val inputLines = input.readLines()

    val cloth = Array(1000) {IntArray(1000) {0} }
    val claims = ArrayList<ClothClaim>()

    // Measure cloth claims
    for (i in 0 until inputLines.size) {

        claims.add(i, ClothClaim(inputLines[i]))

        for (x in claims[i].sideDistance until claims[i].sideDistance + claims[i].width) {
            for (y in claims[i].topDistance until claims[i].topDistance + claims[i].height) {
                cloth[x][y]++
            }
        }
    }

    for (i in 0 until claims.size) {
        for (x in claims[i].sideDistance until claims[i].sideDistance + claims[i].width) {
            for (y in claims[i].topDistance until claims[i].topDistance + claims[i].height) {
                if (cloth[x][y] > 1) claims[i].overlapping = true
            }
        }
    }

    var multiClaims = 0

    for (x in 0 until cloth.size) {
        for (y in 0 until cloth[x].size) {
            if (cloth[x][y] > 1) multiClaims++
        }
    }

    System.out.println(multiClaims)

    claims.forEach { if (!it.overlapping) System.out.println(it.id)}
}