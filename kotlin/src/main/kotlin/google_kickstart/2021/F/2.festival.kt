import java.util.*

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (D, N, K) = input_ints()
        val st = MutableList(D + 1) { mutableListOf<Int>() }
        val ed = MutableList(D + 1) { mutableListOf<Int>() }
        repeat(N) {
            val (h, s, e) = input_ints()
            st[s].add(h)
            ed[e].add(h)
        }

        var ret = 0L
        var tk = 0
        val top = PriorityQueue<Int>()
        val removed_top = hashMapOf<Int, Int>()
        val remain = PriorityQueue<Int>(compareBy { -it })
        val removed_remain = hashMapOf<Int, Int>()
        var top_sum = 0L

        fun add(h: Int) {
            top.add(h)
            top_sum += h
            tk += 1
            if (tk > K) {
                while (removed_top.getOrDefault(top.peek(), 0) > 0) {
                    removed_top[top.peek()] = removed_top[top.peek()]!! - 1
                    top.poll()
                }
                val x = top.poll()
                top_sum -= x
                remain.add(x)
                tk -= 1
            }
        }

        fun remove(h: Int) {
            while (remain.size > 0 && removed_remain.getOrDefault(remain.peek(), 0) > 0) {
                removed_remain[remain.peek()] = removed_remain[remain.peek()]!! - 1
                remain.poll()
            }

            if (remain.size > 0 && remain.peek() >= h) {
                removed_remain[h] = removed_remain.getOrDefault(h, 0) + 1
            } else {
                removed_top[h] = removed_top.getOrDefault(h, 0) + 1
                top_sum -= h
                tk -= 1
                if (remain.size > 0) {
                    val x = remain.poll()
                    top.add(x)
                    top_sum += x
                    tk += 1
                }
            }
        }

        for (d in 1..D) {
            for (h in st[d])
                add(h)
            if (top_sum > ret) {
                ret = top_sum
            }
            for (h in ed[d]) {
                remove(h)
            }
        }
        println("Case #${it}: $ret")
    }
}
