fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M, P) = input_ints()
        val count = MutableList(P) { 0 }
        repeat(N) {
            for ((i, ch) in readLine()!!.withIndex())
                if (ch == '1')
                    count[i] += 1
        }
        val forbid = (1..M).map { readLine()!!.toList() }.toHashSet()
        fun testset2(): Int {
            val low = (0 until P).map { minOf(count[it], N - count[it]) }
            val k = minOf(7, P)
            if (low.map { if (it == (N + 1) / 2) 1 else 0 }.sum() >= 7) return low.sum()

            val subi = (0 until P).sortedBy { -low[it] }.subList(0, k)
            val subs = subi.toHashSet()

            val rem_sum = (0 until P).map { if (it in subs) 0 else low[it] }.sum()
            var min_sub = Int.MAX_VALUE
            var cur = MutableList(P) { i -> if (i in subs) '.' else if (low[i] == count[i]) '0' else '1' }
//            println(cur)
            for (bits in 0 until (1 shl k)) {
                var s = 0
                for (i in 0 until k)
                    if (bits and (1 shl i) > 0) {
                        s += count[subi[i]]
                        cur[subi[i]] = '0'
                    } else {
                        s += N - count[subi[i]]
                        cur[subi[i]] = '1'
                    }
                if (cur !in forbid) {
//                    println("$cur : $s")
                    min_sub = minOf(min_sub, s)
                }
            }
            return rem_sum + min_sub
        }

        fun testset3(): Int {
            val L = 101
            var pre = listOf(0 to listOf<Int>())
            val fb = forbid.map { it.map { if (it == '1') 1 else 0 } }.toHashSet()
            val cnt = (0 until P).map { listOf(count[it], N - count[it]) }
            for (i in 0 until P) {
                var cur = mutableListOf<Pair<Int, List<Int>>>()
                for ((s, p) in pre) {
                    for (x in 0..1) {
                        cur.add(s + cnt[i][x] to p + x)
                    }
                }
                cur.sortBy { it.first }
                pre = cur.subList(0, minOf(L, cur.size))
            }
            for ((s, p) in pre)
                if (p !in fb)
                    return s
            return -1
        }
        println("Case #${it}: ${testset3()}")
    }
}
