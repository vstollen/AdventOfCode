package frequency

import java.io.File

fun main(args : Array<String>) {

    val input = File("input/day1_frequency.txt")

    var frequency = 0

    input.forEachLine { line ->
        if (line[0] == '+') {
            frequency += line.filter { it != '+' }.toInt()
        } else {
            frequency -= line.filter { it != '-' }.toInt()
        }
    }

    System.out.println(frequency)
}