package frequency

import java.io.File

fun main(args : Array<String>) {

    val input = File("input/day1_frequency.txt")

    var frequency = 0

    val reachedFrequencys = HashSet<Int>()

    reachedFrequencys.add(frequency)

    var reachedFrequencyTwice = false

    while (!reachedFrequencyTwice) {
        input.forEachLine { line ->

            if (reachedFrequencyTwice) return@forEachLine

            if (line[0] == '+') {
                frequency += line.filter { it != '+' }.toInt()
            } else {
                frequency -= line.filter { it != '-' }.toInt()
            }

            if (!reachedFrequencys.add(frequency)) {
                System.out.println(frequency)
                reachedFrequencyTwice = true
                return@forEachLine
            }
        }
    }
}