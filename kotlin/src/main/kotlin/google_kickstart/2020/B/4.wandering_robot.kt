package google_kickstart.`2020`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val line = input_ints()
        val (W, H, L, U, R) = line
        val D = line.last()

        val prob = DoubleArray(W + 1)

        fun valid(y: Int, x: Int) = !(x in L..R && y in U..D)

        prob[W] = if (D == H && R == W) 0.0 else 1.0
        for (x in W - 1 downTo 1) prob[x] = if (valid(H, x)) prob[x + 1] else .0

        for (y in H - 1 downTo 1) {
            if (!valid(y, W)) prob[W] = .0
            for (x in W - 1 downTo 1)
                prob[x] = if (valid(y, x)) 0.5 * prob[x + 1] + 0.5 * prob[x] else .0
        }
        println("Case #${it}: ${prob[1]}")
    }
}
// todo: test set 2
