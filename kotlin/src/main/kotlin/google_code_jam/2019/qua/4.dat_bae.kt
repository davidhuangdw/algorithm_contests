package google_code_jam.`2019`.qua

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun upperPow(n: Int): Int {
        val k = (0..n).firstOrNull { (1 shl it) >= n }!!
        return 1 shl k
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, B, F) = input_ints()
        fun set1(): String {
            var unit = upperPow(N) / 2
            var sz = listOf(N - B)
            while (unit > 0) {
                var cur = (0 until N).map { it / unit % 2 }.joinToString("")
                println(cur)
                val (resp) = split_input()
                var acc = 0
                sz = sz.flatMap { k ->
                    val zeros = (acc until acc + k).map { if (resp[it] == '0') 1 else 0 }.sum()
                    acc += k
                    listOf(zeros, k - zeros)
                }
                unit /= 2
            }

            val res = (0 until N).filter { sz[it] == 0 }.joinToString(" ")
            return res
        }

        fun set2(): String {
            val v = MutableList(N - B) { 0 }
            for (i in 0 until 5) {
                val test = (0 until N).map { if (it and (1 shl i) > 0) 1 else 0 }.joinToString("")
                println(test)
                val (resp) = split_input()
                for ((j, ch) in resp.withIndex()) {
                    if (ch == '1') {
                        v[j] += (1 shl i)
                    }
                }
            }

            val res = mutableListOf<Int>()
            var i = 0
            for (x in 0 until N) {
                if (i < v.size && x % 32 == v[i]) {  // bug: i < v.size; bug: x % 32
                    i++
                } else {
                    res.add(x)
                }
            }
            return res.joinToString(" ")
        }
//        val res = set1()
        val res = set2()
        println(res)
        val (resp) = split_input()
        if (resp != "1") {
            throw RuntimeException("failed: $N $B $res: $resp")
        }
    }
}
