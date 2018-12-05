package checksum

import java.io.File

fun main(args : Array<String>) {

    val file = File("input/day2_checksum.txt")

    var letterTwoTimes = 0
    var letterThreeTimes = 0

    file.forEachLine { boxId ->
        val letterMap = HashMap<Char, Int>()
        boxId.forEach { letterMap[it] = 1 + (letterMap[it] ?: 0) }

        if (letterMap.containsValue(2)) letterTwoTimes++
        if (letterMap.containsValue(3)) letterThreeTimes++
    }

    System.out.println(letterTwoTimes * letterThreeTimes)
}