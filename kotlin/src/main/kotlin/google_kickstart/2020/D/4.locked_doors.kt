package google_kickstart.`2020`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val D = mutableListOf(Int.MAX_VALUE)
        for (v in split_input())
            D.add(v.toInt())
        D.add(Int.MAX_VALUE)

        val res = (1..Q).map {
            val (S, K) = input_ints()
            var i = S
            var (l, r) = i to i
            repeat(K - 1) {
                if (D[l - 1] < D[r]) {
                    l -= 1
                    i = l
                } else {
                    r += 1
                    i = r
                }
            }
            i
        }.joinToString(" ")
        println("Case #${it}: $res")
    }
}
// todo: test set 2
