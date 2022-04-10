package google_code_jam.`2022`.qua

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C) = input_ints()
        val res = buildString {
            // first row
            append("..+")
            (2..C).forEach { _ -> append("-+") }
            append("\n")
            for (i in (1..R)) {
                append(if (i == 1) "." else "|")
                (1..C).forEach { _ -> append(".|") }
                append("\n+")
                (1..C).forEach { _ -> append("-+") }
                append("\n")
            }
        }
        print("Case #${it}: \n$res")
    }
}
