package google_code_jam.`2022`.`1a`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (S) = split_input()

        val res = buildString {
            var pre = 'Z' + 1
            var pi = 0
            var i = 0
            while (i < S.length) {
                if (pre < S[i]) {
                    append(S.substring(pi, i))
                }
                pre = S[i]
                var j = i
                while (j < S.length && S[j] == pre) {
                    append(pre)
                    j++
                }
                pi = i
                i = j
            }
        }
        println("Case #${it}: $res")
    }
}
