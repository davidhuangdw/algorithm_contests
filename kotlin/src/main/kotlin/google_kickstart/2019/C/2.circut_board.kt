package google_kickstart.`2019`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C, K) = input_ints()
        val V = (1..R).map { input_ints() }
        val longest = List(R) { MutableList(C) { 0 } }

        for ((r, row) in V.withIndex()) {
            val lows = ArrayDeque<Int>()
            val highs = ArrayDeque<Int>()
            var j = 0
            for (i in 0 until C) {
                while (j < C) {
                    while (lows.isNotEmpty() && row[lows.last()] >= row[j])
                        lows.removeLast()
                    lows.add(j)
                    while (highs.isNotEmpty() && row[highs.last()] <= row[j])
                        highs.removeLast()
                    highs.add(j)
                    if (row[highs.first()] - row[lows.first()] > K)
                        break
                    j += 1
                }
                longest[r][i] = j - i

                if (lows.isNotEmpty() && lows.first() == i)
                    lows.removeFirst()
                if (highs.isNotEmpty() && highs.first() == i)
                    highs.removeFirst()
            }
        }

        var res = 0
        for (j in 0 until C) {
            val suf = MutableList(R) { 0 }
            val st = mutableListOf<Int>()
            for (i in R - 1 downTo 0) {
                while (st.isNotEmpty() && longest[st.last()][j] >= longest[i][j])
                    st.removeLast()
                suf[i] = if (st.isEmpty()) R else st.last()
                st.add(i)
            }
            st.clear()
            for (i in 0 until R) {
                while (st.isNotEmpty() && longest[st.last()][j] >= longest[i][j])
                    st.removeLast()
                val pre = if (st.isEmpty()) -1 else st.last()
                res = maxOf(res, longest[i][j] * (suf[i] - pre - 1))
                st.add(i)
            }
        }

        println("Case #${it}: $res")
    }
}
