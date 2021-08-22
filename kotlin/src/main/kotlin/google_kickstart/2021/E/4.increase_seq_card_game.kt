fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (NN) = input_longs()
        fun testset2(N: Int): Double {
            val f = MutableList(N + 1) { 0.0 }
            val pre = MutableList(N + 1) { 0.0 }
            for (x in 1..N) {
                f[x] = 1 + pre[x - 1] / x
                pre[x] = pre[x - 1] + f[x]
            }
            return f[N]
        }

        fun testset3(N: Long): Double {
            val gamma = 0.57721566490153286061
            return if (N <= 1000_000L)
                testset2(N.toInt())
            else
                kotlin.math.ln(N.toDouble()) + gamma
        }
//        for(x in 1000_001L..1000_100L){
//            val a = testset2(x.toInt())
//            val b = testset3(x)
//            println("$x $a $b ${a-b}")
//        }
        println("Case #${it}: ${testset3(NN)}")
    }
}
