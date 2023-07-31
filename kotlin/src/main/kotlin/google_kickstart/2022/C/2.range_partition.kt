package google_kickstart.`2022`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (N, X, Y) = input_ints()
        fun set2(): String {
            val IMP = "IMPOSSIBLE\n"
            val s = (1..N).sum()
            if (s % (X + Y) != 0) return IMP
            var tar = s / (X + Y) * X
            val pick = mutableListOf<Int>()
            for (i in N downTo 1) {
                if (tar >= i) {
                    tar -= i
                    pick.add(i)
                }
            }
            if (tar != 0) return IMP
            return buildString {
                appendLine("POSSIBLE")
                appendLine(pick.size)
                appendLine(pick.reversed().joinToString(" "))
            }
        }

        val r = set2()
        print("Case #${it}: $r")
    }
}
