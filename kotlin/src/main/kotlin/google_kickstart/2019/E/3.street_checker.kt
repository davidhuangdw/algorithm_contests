package google_kickstart.`2019`.E

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (L, R) = input_ints()
        fun testset1(): Int {
            fun calc(x: Int) = if (x % 2 == 1) 1 else -1
            val diff = MutableList(R - L + 1) { calc(it + L) + if (it + L == 1) 0 else 1 }
            var x = 2
            while (x * x <= R) {
                for (v in maxOf(L, x * x)..R) {
                    if (v % x == 0)
                        diff[v - L] += calc(x) + if (v / x == x) 0 else calc(v / x)
                }
                x += 1
            }
            var cnt = 0
            for (d in diff)
                if (d in -2..2)
                    cnt += 1
            return cnt
        }
        // todo: testset2
        println("Case #${it}: ${testset1()}")
    }
}
