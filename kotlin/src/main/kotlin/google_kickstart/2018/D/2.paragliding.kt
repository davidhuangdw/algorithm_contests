fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        fun build_series(line: List<Int>, n: Int): List<Long> {
            var (x, y) = line.subList(0, 2).map { it.toLong() }
            val (a, b, c, m) = line.subList(2, line.size)
            val series = mutableListOf(x, y)
            repeat(n - 2) {
                x = y.also { y = (a * y + b * x + c) % m + 1 }
                series.add(y)
            }
            return series
        }

        data class Node(val x: Long, val y: Long)

        val towers = mutableListOf<Node>()
        val balloons = mutableListOf<Node>()
        for ((x, y) in build_series(input_ints(), N).zip(build_series(input_ints(), N)))
            towers.add(Node(x, y))
        for ((x, y) in build_series(input_ints(), K).zip(build_series(input_ints(), K)))
            balloons.add(Node(x, y))

        fun testset1(): Int {
            var cnt = 0
            for ((bx, by) in balloons) {
                for ((tx, ty) in towers) {
                    if (ty - by >= kotlin.math.abs(tx - bx)) {
                        cnt += 1
                        break
                    }
                }
            }
            return cnt
        }

        fun testset2(): Int {
            towers.sortBy { it.x }      // bug: sortBy, not sortedBy
            balloons.sortBy { it.x }

            // left
            val done = hashSetOf<Int>()
            var j = 0
            var mx = Long.MIN_VALUE
            for ((i, b) in balloons.withIndex()) {
                val (x, y) = b
                while (j < N && towers[j].x <= x) {
                    mx = maxOf(mx, towers[j].x + towers[j].y)
                    j += 1
                }
                if (mx >= x + y)
                    done.add(i)
            }

            j = N - 1
            mx = Long.MIN_VALUE             // bug: MIN_VALUE, not 0L
            for (i in balloons.size - 1 downTo 0) {
                if (i !in done) {
                    val (x, y) = balloons[i]
                    while (j >= 0 && towers[j].x >= x) {
                        mx = maxOf(mx, towers[j].y - towers[j].x)
                        j -= 1
                    }
                    if (mx >= y - x)
                        done.add(i)
                }
            }
            return done.size
        }
        println("Case #${it}: ${testset2()}")
    }
}
