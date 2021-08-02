package google_kickstart.`2019`.D

import java.util.*
import kotlin.math.absoluteValue

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (K, N) = input_ints()
        val X = input_longs()
        val C = input_longs()

        fun testset1(): Long {
            var min = Long.MAX_VALUE
            for (w in 0 until N) {
                val costs = mutableListOf<Long>()
                for (i in 0 until N)
                    if (i != w)
                        costs.add(C[i] + (X[i] - X[w]).absoluteValue)
                costs.sort()
                min = minOf(min, C[w] + costs.subList(0, K).sum())
            }
            return min
        }

        fun testset2(): Long {
            val (A, B) = K/2 to (K+1)/2     // divide K into K/2 & (K+1)/2
            data class Node(val x: Long, val c: Long)
            val nodes = X.zip(C).map{Node(it.first, it.second)}.sortedBy{it.x}

            val left = MutableList(N){0L}
            val sub = PriorityQueue<Long>(compareBy { -it })
            var sum = 0L
            for(i in 0 until A){
                val v = nodes[i].c - nodes[i].x
                sub.add(v)
                sum += v
            }
            for(i in A until N - B){
                left[i] = sum
                val v = nodes[i].c - nodes[i].x
                sub.add(v)
                sum += v
                sum -= sub.poll()
            }

            sub.clear()
            val right = MutableList(N){0L}
            sum = 0L
            for(i in N-1 downTo N-B){
                val v = nodes[i].c + nodes[i].x
                sub.add(v)
                sum += v
            }
            for(i in N-B-1 downTo A){
                right[i] = sum
                val v = nodes[i].c + nodes[i].x
                sub.add(v)
                sum += v
                sum -= sub.poll()
            }

            var min = Long.MAX_VALUE
            for(i in A until N-B){
                min = minOf(min, left[i] + right[i] + nodes[i].c + (A-B)*nodes[i].x)
            }
            return min
        }
        println("Case #${it}: ${testset2()}")
    }
}
