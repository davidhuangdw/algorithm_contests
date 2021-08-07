package google_kickstart.`2019`.E

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        data class Node(val c: Long, val e: Long)
        val (D, S) = input_ints()
        val all = (1..S).map {
            val (c, e) = input_longs()
            Node(c, e)
        }.sortedWith { a, b -> compareValues(a.e * b.c, b.e * a.c) }

        val pre_c = all.scan(0L) { acc, v -> acc + v.c }
        val pre_e = all.scan(0L) { acc, v -> acc + v.e }

        fun valid(C: Long, E: Long): Boolean {
            var (l, r) = 1 to S
            while (l <= r) {
                val md = (l + r) / 2
                if (pre_c[md] < C)
                    l = md + 1
                else
                    r = md - 1
            }
            val i = l - 1 // pick partial
            if (i >= S) return false
            return (C - pre_c[l - 1]) * all[i].e <= (pre_e.last() - E - pre_e[l - 1]) * all[i].c
        }

        val res = (1..D).map {
            val (C, E) = input_longs()
            if (valid(C, E)) 'Y' else 'N'
        }.joinToString("")

        println("Case #${it}: $res")
    }
}
