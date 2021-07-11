package google_kickstart.`2021`.D

import java.lang.RuntimeException
import java.util.*
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}


//    val T = readLine()!!.toInt()
    val T = 500
    (1..T).forEach {
        val (N, M) = 4 to 100
        val PS = (1..20).shuffled().subList(0, 8).sorted()
        val lines = (0 until 8 step 2).map{PS[it] to PS[it+1]}
        val tmp = (1..4).map{(1..20).random()}
        val SS = tmp // + tmp
//        val lines = listOf(2 to 3, 5 to 7, 8 to 10, 13 to 20)
//        val SS = listOf(16, 14, 16, 16)

        fun testset1(case: Int): MutableList<Int> {
//            val (N, M) = input_ints()
            val points = hashSetOf<Int>()

            repeat(N){
//                val (l, r) = input_ints()
                val (l, r) = lines[it]
                for(p in l..r)
                    points.add(p)
            }

//            val S = input_ints()
            val S = SS
            val res = mutableListOf<Int>()
            for(s in S){
                var x = -1
                if(s in points){
                    x = s
                }else{
                    val l = (s downTo 1).firstOrNull{it in points} ?: -10000
                    val r = (s..1000).firstOrNull{it in points} ?: 10000
                    if(s - l <= r-s) x = l
                    else x = r
                }

                points.remove(x)
                res.add(x)
            }

//            println("Case #${case}: ${res.joinToString(" ")}")
            return res
        }

        fun testset2(case: Int): MutableList<Long> {
//            val (N, M) = input_ints()
            val points = mutableListOf<Long>()
            val lefts = hashSetOf<Long>()
            val rights = hashSetOf<Long>()

            repeat(N){
//                val (l, r) = input_longs()
                val (l, r) = lines[it].toList().map{it.toLong()}
                points.add(l)
                if(l != r)
                    points.add(r)
                lefts.add(l)
                rights.add(r)
            }

            val ord = points.toSortedSet() as TreeSet

            fun removeRightPoint(x: Long){
                ord.remove(x)
                rights.remove(x)
                if(x in lefts)
                    lefts.remove(x)
                else{
                    //insert ,x-1]
                    ord.add(x-1)
                    rights.add(x-1)
                }
            }

            fun removeLeftPoint(x: Long){
                ord.remove(x)
                lefts.remove(x)
                if(x in rights)
                    rights.remove(x)
                else{
                    //insert [x+1,
                    ord.add(x+1)
                    lefts.add(x+1)
                }
            }

            fun insertLeftPoint(x: Long){
                assert(x !in lefts)
                if(x !in rights)
                    ord.add(x)
                lefts.add(x)
            }

            fun insertRightPoint(x: Long){
                assert(x !in rights)
                if(x !in lefts)
                    ord.add(x)
                rights.add(x)
            }
            val res = mutableListOf<Long>()
//            val S = input_longs()
            val S = SS.map{it.toLong()}
            for(s in S){
                println("-".repeat(100))
                println(ord)
                println(lefts)
                println(rights)
                var x = -1L
                if(s in lefts){
                    x = s
                    removeLeftPoint(x)
                }else if(s in rights){
                    x = s
                    removeRightPoint(x)
                }else{
                    val l = ord.lower(s)
                    val r = ord.higher(s)
                    if(l != null && (r== null || s - l <= r - s)){
                        if(l in rights){
                            removeRightPoint(l)
                            x = l
                        }else{
                            x = s
                            insertRightPoint(x-1)
                            insertLeftPoint(x+1)
                        }
                    }else{
                        if(r in lefts){
                            x = r
                            removeLeftPoint(r)
                        }else{
                            x = s
                            insertRightPoint(x-1)
                            insertLeftPoint(x+1)
                        }
                    }
                }
                res.add(x)
            }
            return res

//            println("Case #${case}: ${res.joinToString(" ")}")
        }

        val a = testset1(it)
        val b = testset2(it)

        if(!a.zip(b).all{it.first.toLong() == it.second}){
            println(lines)
            println(SS)
            println(a)
            println(b)
            throw RuntimeException("wrong!")
        }
//        testset1(it)
    }
}
