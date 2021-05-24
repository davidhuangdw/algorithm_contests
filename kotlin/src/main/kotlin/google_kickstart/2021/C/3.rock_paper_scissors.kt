package google_kickstart.`2021`.C
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    val S = "RSP"
    val (X,) = input_ints()
    (1..T).forEach {
        val (W, E) = input_ints()
//        val count = IntArray(3){0}
//        val line = IntArray(60){0}
//        count[0] = 1
//        for(i in 1 until 60){
//            var ma = (0..2).maxByOrNull { k -> count[(k+2)% 3]*W + count[k] * E }!!
//            line[i] = ma
//            count[ma] += 1
//        }
//        println("Case #${it}: ${line.map{S[it]}.joinToString("")}")

        val best = hashMapOf<Triple<Int, Int, Int>, Pair<Double, Char>>()
        for(r in 60 downTo 0)
            for(s in 60-r downTo 0)
                for(p in 60-r-s downTo 0){
                    val sum = r + s + p
                    val cur = Triple(r, s, p)
                    best[cur] = 0.0 to 'R'
                    if(sum == 60)
                        continue
                    val p_s = if(sum == 0) 1.0/3 else 1.0*p/sum
                    val p_r = if(sum == 0) 1.0/3 else 1.0*s/sum
                    val p_p = if(sum == 0) 1.0/3 else 1.0*r/sum
                    // choose r:
                    if(r+1 <= 60){
                        val v = best[Triple(r+1, s, p)]!!.first + W * p_s + E * p_r
                        if(best[cur]!!.first < v)
                            best[cur] = v to 'R'
                    }
                    if(s+1 <= 60){
                        val v = best[Triple(r, s+1, p)]!!.first + W * p_p + E * p_s
                        if(best[cur]!!.first < v)
                            best[cur] = v to 'S'
                    }
                    if(p+1 <= 60){
                        val v = best[Triple(r, s, p+1)]!!.first + W * p_r + E * p_p
                        if(best[cur]!!.first < v)
                            best[cur] = v to 'P'
                    }
                }
        var (r, s, p) = listOf(0, 0, 0)
//        println(best[Triple(0, 0, 0)])
        val line = buildString {
            repeat(60){
                val ch = best[Triple(r,s, p)]!!.second
                append(ch)
                if(ch == 'R') r+=1
                else if(ch == 'S') s+=1
                else p+=1
            }
        }
        println("Case #${it}: $line")
    }
}
