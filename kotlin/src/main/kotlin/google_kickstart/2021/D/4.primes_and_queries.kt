package google_kickstart.`2021`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    class FenwickTree(val N: Int) {
        private val nodes = LongArray(N + 1)
        fun build(values: List<Long>) {
            assert(values.size == N)
            val pre = values.scan(0L, Long::plus)
            nodes[0] = 0L
            for (i in 1..N)
                nodes[i] = pre[i] - pre[i and i - 1]
        }

        fun add(ii: Int, x: Long) {
            var i = ii
            while (i <= N) {
                nodes[i] = nodes[i] + x
                i += i and -i
            }
        }

        fun set(ii: Int, x: Long) {
            val diff = x - (preSum(ii) - preSum(ii - 1))
            add(ii, diff)
        }

        fun preSum(ii: Int): Long {
            var i = ii
            var s = 0L
            while (i > 0) {
                s += nodes[i]
                i -= i and -i
            }
            return s
        }
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q, PP) = input_ints()
        val P = PP.toLong()
        val A = input_longs()
        val S = 4

        fun calc(v: Long, s: Int): Long {
            var x = 1L
            repeat(s) { x *= v }
            var y = 1L
            repeat(s) { y *= (v % P) }
            x -= y
            if (x == 0L) return 0L
            var cnt = 0L
            while (x % P == 0L) {
                cnt += 1
                x /= P
            }
            return cnt
        }

        val tr = (0 until 4).map { s ->
            val t = FenwickTree(N)
            if (it > 0)
                t.build(A.map { calc(it, s + 1) })
            t
        }

        val res = mutableListOf<Long>()
        repeat(Q) {
            val line = input_ints()
            if (line[0] == 1) {
                val (_, i, v) = line
                repeat(4) { s ->
                    tr[s].set(i, calc(v.toLong(), s + 1))
                }
            } else {
                val (_, s, l, r) = line
                res.add(tr[s - 1].preSum(r) - tr[s - 1].preSum(l - 1))
            }
        }
        println("Case #${it}: ${res.joinToString(" ")}")
    }
}
