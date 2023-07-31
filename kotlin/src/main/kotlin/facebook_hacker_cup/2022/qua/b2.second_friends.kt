fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C) = input_ints()
        val V = (1..R).map { split_input()[0].toMutableList() }
        fun valid(i: Int, j: Int) = i in 0 until R && j in 0 until C
        fun around(i: Int, j: Int): List<Pair<Int, Int>> = listOf(i + 1 to j, i - 1 to j, i to j - 1, i to j + 1)
        fun deg(ci: Int, cj: Int): Int {
            var d = 0
            for ((i, j) in around(ci, cj)) {
                if (valid(i, j) && V[i][j] !in "X#") {
                    d++
                }
            }
            return d
        }

        fun current(): String = V.map { it.joinToString("") }.joinToString("\n")
        fun solve(): String {
            if (V.all { it.all { it != '^' } }) {
                return "Possible\n" + current()
            }
            val que = ArrayDeque<Pair<Int, Int>>()
            for (i in 0 until R)
                for (j in 0 until C) {
                    val ch = V[i][j]
                    if (ch !in "X#" && deg(i, j) < 2) {
                        if (ch == '^')
                            return "Impossible"
                        V[i][j] = 'X'
                        que.add(i to j)
                    }
                }
            while (que.isNotEmpty()) {
                val (ci, cj) = que.removeFirst()
                for ((i, j) in around(ci, cj)) {
                    if (valid(i, j) && V[i][j] !in "X#" && deg(i, j) < 2) {
                        val ch = V[i][j]
                        if (ch == '^')
                            return "Impossible"
                        V[i][j] = 'X'
                        que.add(i to j)
                    }
                }
            }

            for (i in 0 until R)
                for (j in 0 until C)
                    if (V[i][j] == 'X')
                        V[i][j] = '.'
                    else if (V[i][j] == '.')
                        V[i][j] = '^'

            return "Possible\n" + current()
        }
        println("Case #${it}: ${solve()}")
    }
}
