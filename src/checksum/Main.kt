package checksum

import java.io.File

fun main(args : Array<String>) {

    val file = File("input/day2_checksum.txt")

    // Generate Checksum
    var letterTwoTimes = 0
    var letterThreeTimes = 0

    file.forEachLine { boxId ->
        val letterMap = HashMap<Char, Int>()
        boxId.forEach { letterMap[it] = 1 + (letterMap[it] ?: 0) }

        if (letterMap.containsValue(2)) letterTwoTimes++
        if (letterMap.containsValue(3)) letterThreeTimes++
    }

    System.out.println(letterTwoTimes * letterThreeTimes)

    // Find fabric boxes

    file.forEachLine { boxId ->

        file.forEachLine {
            var difference = 0
            val differentIndices = HashSet<Int>()

            for (i in 0 until it.length) {
                if (boxId[i] != it[i]) {
                    difference++
                    differentIndices.add(i)
                }

                if (difference > 1) break
            }

            if (difference == 1) {
                System.out.println(boxId.removeRange(differentIndices.first(), differentIndices.first() + 1))
            }
        }
    }
}