package google_kickstart.`2020`.E

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n, a, b, c) = input_ints()
        val res = if (!(a in c..n && b in c..n && c in 1..n)) "IMPOSSIBLE"
        else if (a + b - c > n || (a to b == 1 to 1 && n > 1)) "IMPOSSIBLE"
        else {
            var arr = MutableList(n) { 1 }
            var (i, j) = a - c to n - 1 - (b - c)
            for (k in 1..(a - c)) arr[i - k] = n - k
            for (k in 1..(b - c)) arr[j + k] = n - k

            val extra = n - (a - c) - (b - c) - c
            if (a - c > 0) i += extra
            else if (b - c > 0) j -= extra
            else {
                arr[i] = n
                i += extra + 1
            }
            for (k in i..j) arr[k] = n
            arr.joinToString(" ")
        }
        println("Case #${it}: $res")
    }
}
