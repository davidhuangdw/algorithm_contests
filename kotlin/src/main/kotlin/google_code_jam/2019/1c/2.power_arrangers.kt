package google_code_jam.`2019`.`1c`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val N = 119
    val M = 5
    val (T, F) = input_ints()
    (1..T).forEach {
        var rem = (0 until N).toMutableList()
        var miss = mutableListOf<Char>()
        for (j in 0 until M - 1) {
            val ind = hashMapOf<Char, MutableList<Int>>()
            for (i in rem) {
                println(i * M + j + 1)
                val ch = split_input()[0][0]
                if (ch !in ind)
                    ind[ch] = mutableListOf()
                ind[ch]!!.add(i)
            }
            val k = ind.keys.minByOrNull { ind[it]!!.size }!!
            miss.add(k)
            rem = ind[k]!!
        }
        miss.add((0 until M).map { 'A' + it }.firstOrNull { it !in miss }!!)
        miss[M - 1] = miss[M - 2].also {
            miss[M - 2] = miss[M - 1]
        }  // bug: the last loop only one in rem(compare to missing char)
        val res = miss.joinToString("")
        println(res)
        val (r) = split_input()
        if (r == "N")
            throw RuntimeException("failed: $miss")
//        println("Case #${it}: ")
    }
}
