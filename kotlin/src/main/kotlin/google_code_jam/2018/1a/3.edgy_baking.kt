package google_code_jam.`2018`.`1a`

import kotlin.math.sqrt

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (N, _P) = input_ints()
        val P = _P.toDouble() + 1e-9
        val A = (1..N).map { input_ints().map { it * 2.0 } }

        var base = A.sumOf { it.sum() }
        fun set1(): Double {
            val (w, h) = A[0]
            val (a, b) = minOf(w, h) to sqrt(w * w + h * h)
            val k = ((P - base) / b).toInt()
            if (k >= N)
                return base + N * b
            base += k * b + a
            if (base <= P)
                return P
            val l = base - k * (b - a)
            return if (l <= P)
                P
            else
                base - a   // bug: read doc, should always <= P
        }

        fun set2(): Double {
            val p = P - base
            var intervals = listOf(.0 to .0)
            fun merge() {
                var (ll, rr) = intervals[0]
                val merged = mutableListOf<Pair<Double, Double>>()
                for ((l, r) in intervals) {
                    if (l > p) break
                    if (l <= rr) {
                        rr = maxOf(rr, r)
                    } else {
                        merged.add(ll to rr)
                        ll = l
                        rr = r
                    }
                }
                merged.add(ll to rr)
                intervals = merged
//                println(intervals.map{(l, r) -> l+base to r+base})
            }
            for ((w, h) in A) {
                val (a, b) = minOf(w, h) to sqrt(w * w + h * h)
                intervals = (intervals + intervals.map { (l, r) -> (l + a to r + b) }).sortedBy { it.first }
                merge()
            }
            return minOf(P, base + intervals.lastOrNull()!!.second)
        }

        val res = set2()
        println("Case #${it}: $res")
    }
}
