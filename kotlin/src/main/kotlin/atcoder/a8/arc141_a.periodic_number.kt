package atcoder.a8

fun main(){
    fun input_ints() = readLine()!!.split(" ").map { it.toInt() }

    val (T, ) = input_ints()
    repeat(T){
        val s = readLine()!!
        val N = s.toLong()
        val k = s.length
        var mx = List(k-1){"9"}.joinToString("").toLong()
        for(l in 1 until k){
            if(k % l == 0){
                var part = s.substring(0 until l)
                for(d in 0 ..1){
                    val p = (part.toLong()-d).toString()
                    var cur = List(k/l){p}.joinToString("").toLong()
                    if(cur <= N){
                        mx = maxOf(mx, cur)
                        break
                    }
                }
            }
        }
        println(mx)
    }
}