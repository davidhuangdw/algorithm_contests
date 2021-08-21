package google_kickstart.`2018`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val low_bit = List(1 shl 15) { x ->
        var i = 0
        while (i < 15 && x and (1 shl i) == 0)
            i += 1
        i
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val M = (1..N).map { input_ints() }

        val es = mutableListOf<Int>()
        fun dfs(bits: Int): Int {
            if (bits == 0)
                return if (es.size >= 3 && es.sum() > es.maxOrNull()!! * 2) 1 else 0  // bug: >, not >=
            val i = low_bit[bits]
            var cur = bits xor (1 shl i)

            var cnt = dfs(cur)
            for (j in i + 1 until N) {
                if (M[i][j] > 0 && bits and (1 shl j) > 0) {
                    es.add(M[i][j])
                    cnt += dfs(cur xor (1 shl j))
                    es.removeLast()
                }
            }
            return cnt
        }
        println("Case #${it}: ${dfs((1 shl N) - 1)}")
    }
}
