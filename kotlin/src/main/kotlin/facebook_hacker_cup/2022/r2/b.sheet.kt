import java.lang.Integer.max
import java.util.PriorityQueue

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val M = 1000_000_007L
    data class Point(val sum: Long, val y: Int): Comparable<Point>{
        override fun compareTo(other: Point): Int {
            return compareValues(sum, other.sum)
        }
    }
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        class Que{
            val que = PriorityQueue<Point>()
            fun add(point: Point){
                que.add(point)
                if(que.size > K)
                    que.poll()
            }
        }
        val pre = hashMapOf<Int, Que>()
        val all = PriorityQueue<Long>()
        val start = hashMapOf<Int, MutableList<Triple<Int, Int, Int>>>()
        val ends = hashSetOf<Int>()
        val max_start = hashMapOf<Int, Int>()
        fun addAll(s: Long){
            all.add(s)
            if(all.size > K)
                all.poll()
        }
        for(i in 1..N){
            val (a, b, x, y) = input_ints()
            if(a !in start)
                start[a] = mutableListOf()
            start[a]!!.add(Triple(b, x, y))
            max_start[a] = max(max_start.getOrDefault(a, 0), x)
            ends.add(b)
        }
        for(ps in start.values){
            for((b, x, y) in ps) {
                if (y >= max_start.getOrDefault(b, 0)) continue
                if (b !in pre)
                    pre[b] = Que()
                pre[b]!!.add(Point(0, y))
            }
        }
        for(a in ends.sorted()){
            if(!(a in pre && a in start)) continue
            for(p in pre[a]!!.que){
                for((b, x, y) in start[a]!!){
                    if(p.y < x){
                        val sum = p.sum + x - p.y
                        addAll(sum)
                        if(b !in pre)
                            pre[b] = Que()
                        pre[b]!!.add(Point(sum, y))
                    }
                }
            }
        }
        var sum = 0L
        for(v in all){
            sum = (sum+v) % M
        }
        println("Case #${it}: $sum")
    }
}
