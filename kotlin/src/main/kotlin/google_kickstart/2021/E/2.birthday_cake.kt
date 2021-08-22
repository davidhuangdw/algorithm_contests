fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C, K) = input_ints()
        val (r1, c1, r2, c2) = input_ints()

        fun testset1(): Int {
            val dr = (r2 - r1) + (if (r1 == 1) 0 else 1) + (if (r2 == R) 0 else 1)
            val dc = (c2 - c1) + (if (c1 == 1) 0 else 1) + (if (c2 == C) 0 else 1)
            return dr * (c2 - c1 + 1) + dc * (r2 - r1 + 1) + minOf(r1 - 1, R - r2, c1 - 1, C - c2)
        }
        // todo: testset2
        println("Case #${it}: ${testset1()}")
    }
}
