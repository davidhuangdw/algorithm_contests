package google_code_jam.`2008`.qua

import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun timeToInt(s: String): Int {
        val (h, m) = s.split(":").map { it.toInt() }
        return h * 60 + m
    }

    data class Trip(val st: Int, val ed: Int, val i: Int) : Comparable<Trip> {
        override fun compareTo(other: Trip): Int {
            return if (st != other.st) st - other.st else ed - other.ed
        }
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (T) = input_ints()
        val N = input_ints()
        val trips = mutableListOf<Trip>()
        for ((i, n) in N.withIndex()) {
            for (j in (1..n)) {
                val (st, ed) = split_input().map { timeToInt(it) }
                trips.add(Trip(st, ed, i))
            }
        }
        val cur = List(2) { PriorityQueue<Int>() }
        val cnt = IntArray(2)
        trips.sort()
        for ((st, ed, i) in trips) {
            if (cur[i].isEmpty() || cur[i].peek() > st) {
                cnt[i]++
                cur[i].add(0)
            }
            cur[i].poll()
            cur[1 - i].add(ed + T)
        }
        val res = cnt.joinToString(" ")
        println("Case #${it}: $res")
    }
}
