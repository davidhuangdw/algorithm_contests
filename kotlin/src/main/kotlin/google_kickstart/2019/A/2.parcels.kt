package google_kickstart.`2019`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun solve1(): Long {
        val (R, C) = input_ints()
        val A = List(R) { readLine()!!.toMutableList() }

        fun calc(): Long {
            val done = hashMapOf<Pair<Int, Int>, Long>()
            val que = ArrayDeque<Pair<Int, Int>>()
            for (i in 0 until R)
                for (j in 0 until C) {
                    if (A[i][j] == '1') {
                        done[i to j] = 0
                        que.add(i to j)
                    }
                }
            while (que.isNotEmpty()) {
                val pos = que.removeFirst()
                val (x, y) = pos
                for ((i, j) in listOf(x to y - 1, x to y + 1, x - 1 to y, x + 1 to y)) {
                    if (i to j !in done && i in 0 until R && j in 0 until C) {
                        done[i to j] = done[pos]!! + 1
                        que.add(i to j)
                    }
                }
            }
//            println(done)
//            for(i in 0 until R)
//                println((0 until C).map{done[i to it]!!.toString()}.joinToString("\t"))
            return done.values.maxOrNull()!!
        }

        var min = calc()
        for (i in 0 until R)
            for (j in 0 until C) {
                if (A[i][j] == '0') {
                    A[i][j] = '1'
//                    for(l in A) println(l)
                    min = minOf(min, calc())
                    A[i][j] = '0'
                }
            }
        return min
    }

    fun solve(): Int {
        val (R, C) = input_ints()
        val A = List(R) { readLine()!!.toMutableList() }

        val dis = hashMapOf<Pair<Int, Int>, Int>()
        val que = ArrayDeque<Pair<Int, Int>>()
        for (i in 0 until R)
            for (j in 0 until C) {
                if (A[i][j] == '1') {
                    dis[i to j] = 0
                    que.add(i to j)
                }
            }
        while (que.isNotEmpty()) {
            val pos = que.removeFirst()
            val (x, y) = pos
            for ((i, j) in listOf(x to y - 1, x to y + 1, x - 1 to y, x + 1 to y)) {
                if (i to j !in dis && i in 0 until R && j in 0 until C) {
                    dis[i to j] = dis[pos]!! + 1
                    que.add(i to j)
                }
            }
        }

        val points = mutableListOf<Pair<Int, Pair<Int, Int>>>()
        for (i in 0 until R)
            for (j in 0 until C)
                points.add(dis[i to j]!! to (i to j))
        points.sortBy { -it.first }
//        println(points)

        // -k <= (x - xi) - (y- yi)  <= k  -->   (xi-yi)-k <= x - y <= (xi-yi)+k
        // -k <= (x-xi) + (y-yi) <= k  -->  (xi+yi)-k <= x+y <= (xi+yi)+k
        fun can(k: Int): Boolean {
            var (minus_l, minus_r) = Int.MIN_VALUE to Int.MAX_VALUE
            var (plus_l, plus_r) = Int.MIN_VALUE to Int.MAX_VALUE
            for ((dis, pos) in points) {
                if (dis <= k) break
                val (x, y) = pos
                minus_l = maxOf(minus_l, x - y - k)
                minus_r = minOf(minus_r, x - y + k)
                plus_l = maxOf(plus_l, x + y - k)
                plus_r = minOf(plus_r, x + y + k)
            }

            for (x in 0 until R)
                for (y in 0 until C)
                    if (x - y in minus_l..minus_r && x + y in plus_l..plus_r)
                        return true
            return false
        }

        var (l, r) = 0 to R - 1 + C - 1
        while (l <= r) {
            val md = (l + r) ushr 1
            if (can(md)) r = md - 1
            else l = md + 1
        }
        return l
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        println("Case #${it}: ${solve()}")
    }
}
