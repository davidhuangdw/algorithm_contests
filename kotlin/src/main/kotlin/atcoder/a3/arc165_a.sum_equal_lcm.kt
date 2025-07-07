package atcoder.a3

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }

    val (T, ) = input_ints()
    val tests = (1..T).map{ input_ints()[0] }
    val N = minOf(tests.maxOrNull()!!, 4e4.toInt())

    val isPr = BooleanArray(N){true}
    val pr = mutableListOf<Int>()
    for(i in 2 until N){
        if(isPr[i]){
            pr.add(i)
            for(j in (i shl 1) until N step i){
                isPr[j] = false
            }
        }
    }

    for(t in tests){
        var rem = t
        var ps = mutableListOf<Long>()
        for(p in pr){
            if(rem == 1) break
            var cur = 1L
            while(rem %p == 0){
                rem /= p
                cur *= p
            }
            if(cur > 1){
                ps.add(cur)
            }
        }
        if(rem > 1){
            ps.add(rem + 0L)
        }
//        println(listOf(t, ps))
        println(if(ps.size >= 2 && ps.sum() <= t) "Yes" else "No")
    }
}