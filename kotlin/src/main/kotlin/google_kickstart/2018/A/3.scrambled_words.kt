package google_kickstart.`2018`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (L) = input_ints()
        val W = split_input()
        val line = split_input()
        val (S1, S2) = line
        val (N, A, B, C, D) = line.subList(2, line.size).map { it.toLong() }

        var (x, y) = S1[0].toInt() to S2[0].toInt()
        val S = buildString {
            append(x.toChar())
            append(y.toChar())
            for (i in 3..N) {
                x = y.also { y = ((A * y + B * x + C) % D).toInt() }
                append('a' + y % 26)
            }
        }

        fun trans(w: String): Pair<String, HashMap<Char, Int>> {
            val freq = hashMapOf<Char, Int>()
            for (c in w) {
                freq[c] = freq.getOrDefault(c, 0) + 1
            }
            return "${w[0]}${w.last()}" to freq
        }


        val count = hashMapOf<Int, HashMap<Pair<String, HashMap<Char, Int>>, HashSet<Int>>>()
        for ((i, w) in W.withIndex()) {
            val x = trans(w)
            val l = w.length
            if (l !in count)
                count[l] = hashMapOf()
            val map = count[l]!!
            if (x !in map)
                map[x] = hashSetOf()
            map[x]!!.add(i)
        }

        var found = hashSetOf<Int>()
        for ((k, map) in count) {
            val freq = hashMapOf<Char, Int>()
            for ((i, ch) in S.withIndex()) {
                freq[ch] = freq.getOrDefault(ch, 0) + 1
                if (i >= k - 1) {
                    var c = S[i - k + 1]
                    found.addAll(map.getOrDefault("$c$ch" to freq, hashSetOf()))
                    if (freq[c] == 1)
                        freq.remove(c)
                    else
                        freq[c] = freq[c]!! - 1
                }
            }
        }

        println("Case #${it}: ${found.size}")
    }
}
