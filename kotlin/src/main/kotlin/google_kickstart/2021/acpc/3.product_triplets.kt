package google_kickstart.`2021`.acpc

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val M = 2 * 1e5.toInt()
        val (N) = input_ints()
        val A = input_ints()
        val count = MutableList(M + 1) { 0 }
        for (v in A)
            count[v] += 1
        var sum = 0L
        val c0 = count[0]
        if (c0 >= 2) {
            // 0 * 0 = 0
            sum += 1L * c0 * (c0 - 1) * (c0 - 2) / 6
            // 0 * non_0 = 0
            sum += 1L * c0 * (c0 - 1) / 2 * (N - c0)
        }
        for (z in 1..M) {
            val c = count[z]
            if (c == 0) continue
            if (z <= 1) {
                // 1*1 = 1
                if (c - 2 > 0)
                    sum += 1L * c * (c - 1) * (c - 2) / 6
            } else {
                // 1*z = z
                if (c - 1 > 0 && count[1] > 0) {
                    sum += 1L * c * (c - 1) / 2 * count[1]
                }
                var x = 2
                while (x * x <= z) {
                    if (count[x] > 0 && z % x == 0 && count[z / x] > 0) {
                        val y = z / x
                        if (x == y)
                            sum += 1L * count[x] * (count[x] - 1) / 2 * c
                        else
                            sum += 1L * count[x] * count[y] * c
                    }
                    x += 1
                }
            }
        }
        println("Case #${it}: $sum")
    }
}
