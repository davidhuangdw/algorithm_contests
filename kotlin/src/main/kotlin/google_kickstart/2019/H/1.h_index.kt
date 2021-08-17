package google_kickstart.`2019`.H

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()

        val tree = MutableList(N + 1) { 0 }
        fun query(x: Int): Int {
            var i = x
            var sum = 0
            while (i > 0) {
                sum += tree[i]
                i -= i and -i
            }
            return sum
        }

        fun add(x: Int) {
            var i = minOf(x, N)
            while (i <= N) {
                tree[i] += 1
                i += i and -i
            }
        }

        val res = mutableListOf<Int>()
        for (a in input_ints()) {
            add(a)
            var n = res.size + 1
            var (l, r) = 1 to n
            while (l <= r) {
                val md = (l + r) / 2
                if (n - query(md - 1) >= md) l = md + 1
                else r = md - 1
            }
            res.add(r)
        }
        println("Case #${it}: ${res.joinToString(" ")}")
    }
}
