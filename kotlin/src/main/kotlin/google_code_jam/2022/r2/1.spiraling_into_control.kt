package google_code_jam.`2022`.r2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (N, K) = input_ints()
        fun set3(): String {
            val path = mutableListOf<Pair<Int, Int>>()
            var rem = N * N - 1 - K
            var i = 1
            var twice = 1
            var d = N
            var save = (N - 1) * 4 - 2
            var j = -1
            while (i < N * N) {
                if (i >= j && rem >= save) {
                    rem -= save
                    j = i + save + 1
                    if (j - i > 1)
                        path.add(i to j)
                }
                i += d
                save -= 2
                twice = 1 - twice
                if (twice == 0)
                    d--
            }
            if (rem != 0) return "IMPOSSIBLE\n"
            return buildString {
//                appendln(path.size)
                for ((i, j) in path) {
//                    appendln("$i $j")
                }
            }
        }

        val r = set3()
        print("Case #${it}: $r")
    }
}
