package google_kickstart.`2019`.H

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val A = input_longs()
        fun testset1(): Boolean {
            var sum = 0
            val hf = (A.sum() / 2).toInt()
            for ((i, a) in A.withIndex())
                sum = (sum + (a.toInt() % 11) * (i + 1)) % 11
            if (sum % 2 != 0)
                sum += 11
            val dp = List(hf + 1) { if (it == 0) hashSetOf(0) else hashSetOf() }
            for (i in 1..9)
                for (k in hf downTo 1) {
                    for (d in 1..minOf(k, A[i - 1].toInt())) {
                        for (v in dp[k - d])
                            dp[k].add((v + d * i) % 11)
                    }
                }
            return sum / 2 in dp[hf]
        }
        //todo: testset2
        println("Case #${it}: ${if (testset1()) "YES" else "NO"}")
    }
}
