fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val S = readLine()!!.trim()
        val s = S.toMutableList()

        fun testset2(): String {
            val n = s.size
            val pos = hashMapOf<Char, MutableList<Int>>()
            for ((i, ch) in S.withIndex()) {
                if (ch !in pos)
                    pos[ch] = mutableListOf()
                pos[ch]!!.add(i)
            }

            val cs = pos.keys.sortedBy { -pos[it]!!.size }
            if (pos[cs.first()]!!.size > s.size / 2)
                return "IMPOSSIBLE"

            val diffs = MutableList(n) { -1 }
            var i = 0
            for (ch in cs) {
                for (p in pos[ch]!!) {
                    diffs[i] = p
                    i += 2
                    if (i >= n)
                        i = 1
                }
            }

            fun swap(x: Int, y: Int) {
                s[x] = s[y].also { s[y] = s[x] }
            }

            for (j in 1 until n) {
                swap(diffs[0], diffs[j])
                diffs[0] = diffs[j].also { diffs[j] = diffs[0] }
            }
            return s.joinToString("")
        }
        println("Case #${it}: ${testset2()}")
    }
}
