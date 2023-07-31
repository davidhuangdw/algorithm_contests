package google_code_jam.`2017`.`1a`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val line = input_ints()
        val (HD, AD, HK, AK) = line
        val (B, D) = line.subList(4, 6)
        fun set1(): String {
            fun getCount(_d: Int, _b: Int): Int {
                var (d, b) = _d to _b
                var cnt = 0
                var (hd, ad, hk, ak) = listOf(HD, AD, HK, AK)
                while (cnt < 600) {
                    cnt++
                    if (d > 0 && hd - (ak - D) > 0) {
                        d--
                        ak -= D
                    } else if (hk - ad > 0 && hd - ak <= 0) {
                        hd = HD
                    } else if (b > 0) {
                        b--
                        ad += B
                    } else {
                        hk -= ad
                        if (hk <= 0)
                            return cnt
                    }
                    hd -= ak
                    if (hd <= 0)
                        break
                }
                return Int.MAX_VALUE
            }

            var min = Int.MAX_VALUE
            for (d in 0..minOf(100, AK))
                for (b in 0..minOf(100, HK - AD)) {
                    min = minOf(min, getCount(d, b))

                }
            return if (min == Int.MAX_VALUE) "IMPOSSIBLE" else min.toString()
        }

        // todo: set2
        val r = set1()
        println("Case #${it}: $r")
    }
}
