package google_kickstart.`2019`.E

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M) = input_ints()
        var remain = N - 1
        val par = MutableList(N + 1) { it }
        fun find(x: Int): Int {
            if (par[x] != x)
                par[x] = find(par[x])
            return par[x]
        }

        fun union(x: Int, y: Int) {
            val (rx, ry) = find(x) to find(y)
            if (rx != ry) {
                remain -= 1
                par[rx] = ry
            }
        }

        repeat(M) {
            val (x, y) = input_ints()
            union(x, y)
        }
        println("Case #${it}: ${N - 1 + remain}")
    }
}
