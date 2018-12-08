package fabric

import java.io.File

fun main(args : Array<String>) {

    val input = File("input/day3_fabric.txt")

    val cloth = Array(1000) {IntArray(1000) {0} }

    // Measure cloth claims
    input.forEachLine {

        val sideDistance = it.substringAfter("@ ").substringBefore(',').toInt()
        val topDistance = it.substringAfter(',').substringBefore(':').toInt()

        val width = it.substringAfter(": ").substringBefore('x').toInt()
        val height = it.substringAfter('x').toInt()

        for (x in sideDistance until sideDistance + width) {
            for (y in topDistance until topDistance + height) {
                cloth[x][y]++
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

}