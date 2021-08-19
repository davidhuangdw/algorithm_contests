package google_kickstart.`2018`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (F, L) = input_longs()
        fun testset1(): Long {
            val MX = 1e6.toInt()
            val count = MutableList(MX + 1) { 0 }
            fun nine(xx: Int): Boolean {
                var x = xx
                var sum = 0
                while (x > 0) {
                    if (x % 10 == 9) return true
                    sum += x % 10
                    x /= 10
                }
                return sum % 9 == 0
            }
            for (x in 9..MX) {
                count[x] = count[x - 1] + if (nine(x)) 1 else 0
            }
            return L - F + 1 - (count[L.toInt()] - count[F.toInt() - 1])
        }

        fun testset2(): Long {
            fun count(x: Long): Long {
                val pow = (1..18).scan(1L) { acc, i -> acc * 10L }
                val dig = (0..18).map { (x / pow[it] % 10).toInt() }
                val dp = List(19) { List(9) { MutableList(2) { 0L } } }
                for (i in 0..18)
                    for (m in 0 until 9)
                        for (lt in 0 until 2) {
                            var cur = 0L
                            if (i == 0) {
                                // t==9
                                if (lt == 1 || dig[i] == 9)
                                    cur += 1L

                                // t != 9 && sum%9==0
                                val t = (9 - m) % 9
                                if (t != 9 && (lt == 1 || t <= dig[i]))
                                    cur += 1L
                            } else {
                                // t == 9
                                if (lt == 1 || dig[i] == 9)
                                    cur += pow[i]

                                // t != 9
                                if (lt == 1) {
                                    for (t in 0..8)
                                        cur += dp[i - 1][(t + m) % 9][1]
                                } else {
                                    for (t in 0..dig[i] - 1)
                                        cur += dp[i - 1][(t + m) % 9][1]
                                    val t = dig[i]
                                    if (t != 9) {
                                        cur += dp[i - 1][(t + m) % 9][0]
                                    }
                                }
                            }
                            dp[i][m][lt] = cur
                        }
                return dp[18][0][0]
            }
//            for(x in 0 until 30)
//                println("$x: ${count(x.toLong())}")
            return L - (F - 1) - (count(L) - count(F - 1))
        }
        println("Case #${it}: ${testset2()}")
    }
}
