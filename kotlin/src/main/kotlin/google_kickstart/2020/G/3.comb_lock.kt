package google_kickstart.`2020`.G

import kotlin.math.absoluteValue

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    fun input_longs() = split_input().map { it.toLong() }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (w, n) = input_ints()
        val a = input_longs().sorted()
//        val a = (1..w).map{(1..n).random().toLong()}.sorted()
//        val a = listOf(1, 1, 4, 4, 5).map{v -> v.toLong()-1}.sorted()
//        var res = Long.MAX_VALUE
//        for(x in a)
//            res = minOf(res, a.sumOf { y -> minOf((x-y).absoluteValue, n - (x-y).absoluteValue)})

        val pre = a.scan(0, Long::plus)
        fun sub(i: Int, j: Int) = if (i < j) pre[j] - pre[i] else pre[w] + pre[j] - pre[i] // sumOf [i, j)
        var r = 0
        var res = Long.MAX_VALUE
        for (i in 0 until w) {
            if(i > 0 && a[i] == a[i-1]) continue
            if(r == i){
                while(r < i+w && a[r % w] == a[i]) r += 1
                r %= w
            }
            while (a[r] != a[i] && (a[r] - a[i] + n) % n <= n / 2) r = (r + 1) % w
            val up = if (r > i) sub(i, r) - (r - i) * a[i]
            else sub(i, r) - (w - i) * a[i] + (n - a[i]) * r
            val down = when {
                r == i -> 0
                r < i -> (i - r) * a[i] - sub(r, i)
                else -> i * a[i] + (w - r) * (a[i] + n) - sub(r, i)
            }
            res = minOf(res, up + down)
        }

        println("Case #${it}: $res")
    }
}
