package atcoder.a11
fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)
//    for(n in 2..6){
//        for(v in 0 until (1 shl n)){
//            var x = v
//            val s = (1..n).map{if(x.also{x /=2} and 1 >0) 'B' else 'A'}.joinToString("")
//            var t = s
////            println("-".repeat(10) + s)
//            repeat(n-1){
//                val ai = (0 until t.length-1).filter{t[it] == 'A'}
//                val bi = (0 until t.length-1).filter{t[it] == 'B'}
//                t = (ai.map{t[it+1]} +bi.map{t[it+1]}).joinToString("")
////                println(t)
//            }
////            if(t == "B"){println(s)}
//        }
//    }

    val (T) = inputInts()
    repeat(T) {
        val (N) = inputInts()
        val S = readLine()!!
        val i = (0 until N).find{S[it] == 'B'} ?: N
        println(if(i < N && (i until N).all{S[it] == 'B'}) 'B' else 'A')
    }
}
