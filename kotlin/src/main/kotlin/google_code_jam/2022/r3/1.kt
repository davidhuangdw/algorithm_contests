package google_code_jam.`2022`.r3

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T, N, K) = input_ints()
//    (1..T).forEach {
    var done = 0
    while (done < T) {
        val box = input_ints()
        var c = 1
        var col = MutableList(N) { -1 }
        val two = mutableListOf<Int>()
        for ((i, x) in box.withIndex()) {
            if (col[i] > 0)
                continue
            var j = i
            val ind = mutableListOf<Int>()
            while (col[j] <= 0) {
                ind.add(j)
                col[j] = c
                j = box[j] - 1
            }
            if (ind.size > 30) {
                c++
                for (k in ind.size / 2 until ind.size)
                    col[ind[k]] = c
            }
            c++
        }
        val ch = MutableList(N + 1) { it }
        for (i in 1 until two.size step 2) {
            ch[two[i]] = two[i - 1]
        }
        for (i in 0 until N) {
            col[i] = ch[col[i]]
        }

        println(col.joinToString(" "))
        val ret = input_ints()[0]
        if (ret == -1) {
            throw RuntimeException("failed")
        }
        if (ret == 1)
            done++
    }

//        println("Case #${it}: ")
//    }
}
