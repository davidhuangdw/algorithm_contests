fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        fun testset2(): Int {     // O(n*logn)
            val A = input_ints().sorted()
            var d = 1
            var c = 0
            for (v in A) {
                if (v >= d) {
                    c += 1
                    if (c % K == 0)
                        d += 1
                }
            }
            return c
        }

        fun testset3(): Int {     // O(n), count_sort
            val count = hashMapOf<Int, Int>()
            for (x in input_ints()) {
                val v = minOf(x, N)
                count[v] = count.getOrDefault(v, 0) + 1
            }
            var (d, c) = 1 to 0
            for (v in 1..N) {
                repeat(count.getOrDefault(v, 0)) {
                    if (v >= d) {
                        c += 1
                        if (c % K == 0)
                            d += 1
                    }
                }
            }
            return c
        }
        println("Case #${it}: ${testset3()}")
    }
}
