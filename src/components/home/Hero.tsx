import React from 'react';

const Hero = () => {
  return (
    <section className="bg-gradient-to-r from-primary to-secondary text-white py-20">
      <div className="container">
        <div className="max-w-3xl mx-auto text-center">
          <h1 className="text-4xl md:text-5xl font-bold mb-6">
            세상의 모든 보험 보장분석이 무료!
          </h1>
          <p className="text-xl mb-8">
            전문가와 1:1로 궁금한 점을 상담하세요
          </p>
          <button className="btn-primary text-lg">
            무료신청 보험료 분석
          </button>
        </div>
      </div>
    </section>
  );
};

export default Hero; 