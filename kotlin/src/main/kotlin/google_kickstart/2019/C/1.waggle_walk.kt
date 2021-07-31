package google_kickstart.`2019`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, R, C, si, sj) = input_ints()
        val cmd = readLine()!!.trim()

        // idea: cache similar to disjoint_set.union
        fun testset2(): Pair<Int, Int> {
            val vis = hashSetOf<Pair<Int, Int>>()
            val dir = listOf(0 to 1, 0 to -1, 1 to 0, -1 to 0)
            val c_to_d = hashMapOf('E' to 0, 'W' to 1, 'S' to 2, 'N' to 3)
            val nxt = List(4) { hashMapOf<Pair<Int, Int>, Pair<Int, Int>>() }
            fun next(st: Pair<Int, Int>, d: Int): Pair<Int, Int> {
                var pos = st.first + dir[d].first to st.second + dir[d].second
                nxt[d][st] = nxt[d].getOrDefault(st, pos)
                if (nxt[d][st] in vis) {
                    nxt[d][st] = next(nxt[d][st]!!, d)
                }
                return nxt[d][st]!!
            }

            var cur = si to sj
            for (c in cmd) {
                vis.add(cur)
                cur = next(cur, c_to_d[c]!!)
            }
            return cur
        }
        val (i, j) = testset2()
        println("Case #${it}: $i $j")
    }
}
