fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (K) = input_ints()
        fun testset1(): String {
            val xs = mutableListOf<Pair<Int, Int>>()
            val ys = mutableListOf<Pair<Int, Int>>()
            repeat(K) {
                val (x1, y1, x2, y2) = input_ints()
                xs.add(x1 to x2)
                ys.add(y1 to y2)
            }
            fun getMinPos(vs: List<Pair<Int, Int>>): Int {
                var min_dis = Long.MAX_VALUE
                var mx = 0
                for (x in -100..100) {
                    var d = 0L
                    for ((a, b) in vs) {
                        if (x < a) {
                            d += a - x
                        } else if (b < x) {
                            d += x - b
                        }
                    }
                    if (d < min_dis) {
                        min_dis = d
                        mx = x
                    }
                }
                return mx
            }
            return "${getMinPos(xs)} ${getMinPos(ys)}"
        }

        fun testset2(): String {
            val xs = mutableListOf<Int>()
            val ys = mutableListOf<Int>()
            repeat(K) {
                val (x1, y1, x2, y2) = input_ints()
                xs.add(x1)
                xs.add(x2)
                ys.add(y1)
                ys.add(y2)
            }
            xs.sort()
            ys.sort()
            return "${xs[K - 1]} ${ys[K - 1]}"
        }
        println("Case #${it}: ${testset2()}")
    }
}
